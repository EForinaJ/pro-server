// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDistribute is the golang structure for table sys_distribute.
type SysDistribute struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	IsCancel   int         `json:"isCancel"   orm:"is_cancel"   description:""` //
	Images     string      `json:"images"     orm:"images"      description:""` //
	Reason     string      `json:"reason"     orm:"reason"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
