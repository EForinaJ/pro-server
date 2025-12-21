// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserLog is the golang structure for table sys_user_log.
type SysUserLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	Mode       int         `json:"mode"       orm:"mode"        description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
