package dao_permission

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id" dc:"权限ID"`
	PId         int64       `json:"pId" description:"父菜单ID"`
	Permission  string      `json:"permission" dc:"权限标识"`
	Name        string      `json:"name" dc:"权限名称"`
	Description string      `json:"description" dc:"权限描述"`
	CreateTime  *gtime.Time `json:"createTime"  dc:"创建时间"`
}
