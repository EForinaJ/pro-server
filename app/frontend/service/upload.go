package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 定义显示接口
type IUpload interface {
	// 文件上传
	UploadFile(ctx context.Context, file *ghttp.UploadFile) (res []string, err error)
}

// 定义接口变量
var localUpload IUpload

// 定义一个获取接口的方法
func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

// 定义一个接口实现的注册方法
func RegisterUpload(i IUpload) {
	localUpload = i
}
