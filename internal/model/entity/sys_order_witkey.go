// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderWitkey is the golang structure for table sys_order_witkey.
type SysOrderWitkey struct {
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	IsReplaced int         `json:"isReplaced" orm:"is_replaced" description:""` //
	Reason     string      `json:"reason"     orm:"reason"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
