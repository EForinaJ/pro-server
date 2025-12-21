package dao_document

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id         int64       `json:"id"     dc:"文档编号"`
	Title      string      `json:"title"   dc:"文档标题"`
	Content    string      `json:"content"   dc:"文档内容"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
	Status     int         `json:"status"   dc:"文档状态"`
}
