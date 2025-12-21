// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserExperience is the golang structure for table sys_user_experience.
type SysUserExperience struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	Experience int64       `json:"experience" orm:"experience"  description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
