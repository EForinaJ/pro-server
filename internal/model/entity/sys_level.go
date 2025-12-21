// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLevel is the golang structure for table sys_level.
type SysLevel struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	Level      int         `json:"level"      orm:"level"       description:""` //
	Icon       string      `json:"icon"       orm:"icon"        description:""` //
	Experience int         `json:"experience" orm:"experience"  description:""` //
	Benefits   string      `json:"benefits"   orm:"benefits"    description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
