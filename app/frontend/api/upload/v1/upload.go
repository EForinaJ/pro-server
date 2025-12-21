package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/upload" method:"post" tags:"上传" summary:"小文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type UploadFileRes struct {
	Links []string `json:"links" dc:"返回上传的链接"`
}
