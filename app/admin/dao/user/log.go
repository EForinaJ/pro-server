package dao_user

import "github.com/gogf/gf/v2/os/gtime"

type Log struct {
	Name       string      `json:"name" dc:"操作者"`
	CreateTime *gtime.Time `json:"createTime" dc:"操作时间"`
	Mode       int         `json:"mode" dc:"操作内容"`
	Content    string      `json:"content" dc:"详细内容"`
}
