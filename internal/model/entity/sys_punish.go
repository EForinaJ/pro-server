// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPunish is the golang structure for table sys_punish.
type SysPunish struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	Money      float64     `json:"money"      orm:"money"       description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
