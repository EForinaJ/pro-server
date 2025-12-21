package dao_role

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id" dc:"角色ID"`
	Name        string      `json:"name" dc:"角色名称"`
	Code        string      `json:"code" dc:"角色编码"`
	Description string      `json:"description" dc:"角色描述"`
	Status      int         `json:"status" dc:"角色状态 1:禁用,2:启用"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
	Remark      string      `json:"remark" dc:"备注"`
}
