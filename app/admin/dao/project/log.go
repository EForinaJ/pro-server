package dao_project

import "github.com/gogf/gf/v2/os/gtime"

type Log struct {
	Name       string      `json:"name" dc:"操作者"`
	Type       int         `json:"type" dc:"操作类型"`
	CreateTime *gtime.Time `json:"createTime" dc:"操作时间"`
	Content    string      `json:"content" dc:"操作内容"`
}
