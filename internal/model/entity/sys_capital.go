// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCapital is the golang structure for table sys_capital.
type SysCapital struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	Related    string      `json:"related"    orm:"related"     description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	Mode       int         `json:"mode"       orm:"mode"        description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
