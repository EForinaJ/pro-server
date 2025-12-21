// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRefund is the golang structure for table sys_refund.
type SysRefund struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	Money      float64     `json:"money"      orm:"money"       description:""` //
	Mode       int         `json:"mode"       orm:"mode"        description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
