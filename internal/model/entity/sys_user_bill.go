// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserBill is the golang structure for table sys_user_bill.
type SysUserBill struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	RelatedId  int64       `json:"relatedId"  orm:"related_id"  description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	Mode       int         `json:"mode"       orm:"mode"        description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
