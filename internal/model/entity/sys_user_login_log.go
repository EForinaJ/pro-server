// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserLoginLog is the golang structure for table sys_user_login_log.
type SysUserLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	Ip         string      `json:"ip"         orm:"ip"          description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
