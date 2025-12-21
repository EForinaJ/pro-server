package upload

import (
	"context"
	"os"
	dto_upload "server/app/admin/dto/upload"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUpload struct{}

// MergeChunk implements service.IUpload.
func (s *sUpload) MergeChunk(ctx context.Context, req *dto_upload.Merge) (res []string, err error) {
	tmpFolderPath := "./public/chunk/" + req.Identifier
	if !gfile.Exists(tmpFolderPath) {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR)
	}

	tmpFolderFiles, err := os.ReadDir(tmpFolderPath)
	if err != nil {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR)
	}
	fileConfig, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR)
	}
	fjson := gjson.New(fileConfig)
	path := fjson.Get("path").String()

	filePath := "./public/" + path + "/" + gtime.Date() + "/"
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	filePath = filePath + name + gfile.Ext(req.FileName)

	fileHandle, err := gfile.Create(filePath)
	if err != nil {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR)
	}
	defer fileHandle.Close()

	for index, _ := range tmpFolderFiles {
		tmpFileName := gconv.String(index) + ".temp"
		fileBuffer, err := os.ReadFile(tmpFolderPath + "/" + tmpFileName)
		if err != nil {
			return nil, utils_error.Err(response.FILE_SAVE_ERROR)
		}

		_, err = fileHandle.Write(fileBuffer)
		if err != nil {
			return nil, utils_error.Err(response.FILE_SAVE_ERROR)
		}

	}

	// 删除目录
	err = gfile.Remove(tmpFolderPath)
	if err != nil {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR)
	}

	path = filePath[1:]
	mediaType := gfile.Ext(req.FileName)
	mediaType = mediaType[1:]
	entites := &do.SysMedia{
		Path:       path,
		Name:       name + gfile.Ext(req.FileName),
		OrName:     req.FileName,
		Size:       gconv.String(req.Size),
		MediaType:  gstr.ToUpper(mediaType),
		Ext:        gfile.Ext(req.FileName),
		CreateTime: gtime.Now(),
	}

	_, err = dao.SysMedia.Ctx(ctx).Data(&entites).Insert()
	if err != nil {
		return nil, err
	}

	res = []string{
		gconv.String(entites.Path),
	}

	return
}

// UploadChunk implements service.IUpload.
func (s *sUpload) UploadChunk(ctx context.Context, req *dto_upload.Chunk) (err error) {
	folderPath := "./public/chunk/" + req.Identifier
	fileName := folderPath + "/" + req.ChunkNumber + ".temp" //文件完全路径名
	req.File.Filename = fileName

	_, err = req.File.Save(folderPath)
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED)
	}
	return
}

// ChunkIdentifier implements service.IUpload.
func (s *sUpload) ChunkIdentifier(ctx context.Context, identifier string) (res []int, err error) {
	folderPath := "./public/chunk/" + identifier

	if gfile.Exists(folderPath) {
		list, err := gfile.ScanDir(folderPath, "*", true)
		if err != nil {
			return nil, err
		}

		fileNames := make([]int, len(list))
		for i, v := range list {
			tmpNames := gstr.Split(gfile.Basename(v), ".")
			fileNames[i] = gconv.Int(tmpNames[0])
		}
		return fileNames, nil
	}

	return
}

// UploadMinFile implements service.IUpload.
func (s *sUpload) UploadMinFile(ctx context.Context, file *ghttp.UploadFile) (res []string, err error) {

	fileConfig, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, err
	}
	fjson := gjson.New(fileConfig)
	path := fjson.Get("path").String()

	filePath := "./public/" + path + "/" + gtime.Date()
	orName := file.Filename
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	file.Filename = name + gfile.Ext(file.Filename)
	_, err = file.Save(filePath, false)
	if err != nil {
		return nil, err
	}
	entites := &do.SysMedia{
		Path:       filePath[1:] + "/" + file.Filename,
		Ext:        gfile.Ext(file.Filename),
		MediaType:  gfile.Ext(file.Filename)[1:],
		Name:       name + gfile.Ext(file.Filename),
		CreateTime: gtime.Now(),
		OrName:     orName,
		Size:       gconv.String(file.Size),
	}

	_, err = dao.SysMedia.Ctx(ctx).Data(&entites).Insert()
	if err != nil {
		return nil, err
	}

	res = []string{
		gconv.String(entites.Path),
	}
	return
}

func init() {
	service.RegisterUpload(&sUpload{})
}
