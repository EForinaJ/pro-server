// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderWitkeyAcceptLog is the golang structure for table sys_order_witkey_accept_log.
type SysOrderWitkeyAcceptLog struct {
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
