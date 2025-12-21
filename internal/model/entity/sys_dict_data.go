// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	Key        string      `json:"key"        orm:"key"         description:""` //
	Value      string      `json:"value"      orm:"value"       description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
