// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRecharge is the golang structure for table sys_recharge.
type SysRecharge struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	PayMode    int         `json:"payMode"    orm:"pay_mode"    description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
	Remark     string      `json:"remark"     orm:"remark"      description:""` //
}
