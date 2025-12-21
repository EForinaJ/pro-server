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
	Money      float64     `json:"money"      orm:"money"       description:""` //
	PayType    int         `json:"payType"    orm:"pay_type"    description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
	Remark     string      `json:"remark"     orm:"remark"      description:""` //
}
