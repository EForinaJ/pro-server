package dao_media

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id"         dc:"附件ID"`
	NickName   string      `json:"nickName"   dc:"上传的用户"`
	Path       string      `json:"path"       dc:"存放路径"`
	Name       string      `json:"name"       dc:"文件名称"`
	OrName     string      `json:"orName"     dc:"原始文件名称"`
	Size       string      `json:"size"       dc:"文件大小"`
	Ext        string      `json:"ext"        dc:"文件后缀"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
	MediaType  string      `json:"mediaType"  dc:"文件类型"`
	Remark     string      `json:"remark"     dc:"备注"`
}
