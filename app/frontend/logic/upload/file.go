package upload

import (
	"context"
	"server/app/common/consts"
	"server/internal/dao"
	"server/internal/model/do"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// UploadFile implements service.IUpload.
func (s *sUpload) UploadFile(ctx context.Context, file *ghttp.UploadFile) (res []string, err error) {
	fileConfig, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, err
	}
	fjson := gjson.New(fileConfig)
	path := fjson.Get("path").String()

	filePath := "./public/" + path + "/" + gtime.Date()

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
		OrName:     file.Filename,
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
