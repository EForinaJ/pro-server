// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderLog is the golang structure for table sys_order_log.
type SysOrderLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
