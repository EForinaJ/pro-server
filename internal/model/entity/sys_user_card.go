// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserCard is the golang structure for table sys_user_card.
type SysUserCard struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	Number     string      `json:"number"     orm:"number"      description:""` //
	NickName   string      `json:"nickName"   orm:"nick_name"   description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
