// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAftersales is the golang structure for table sys_aftersales.
type SysAftersales struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	Reason     string      `json:"reason"     orm:"reason"      description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
