package dao_manage

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"用户ID"`
	Name       string      `json:"name"  dc:"用户昵称"`
	Phone      string      `json:"phone" dc:"手机号码"`
	Sex        int         `json:"sex" dc:"性别"`
	Avatar     string      `json:"avatar" dc:"头像地址"`
	Status     int         `json:"status" dc:"帐号状态（1停用,2正常）"`
	Roles      []string    `json:"roles" dc:"所拥有的角色"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
	Remark     string      `json:"remark" dc:"备注"`
}
