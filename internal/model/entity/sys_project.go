// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProject is the golang structure for table sys_project.
type SysProject struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	OrderId    int64       `json:"orderId"    orm:"order_id"    description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	Images     string      `json:"images"     orm:"images"      description:""` //
	Money      float64     `json:"money"      orm:"money"       description:""` //
	ServiceFee float64     `json:"serviceFee" orm:"service_fee" description:""` //
	Commission float64     `json:"commission" orm:"commission"  description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	StartTime  *gtime.Time `json:"startTime"  orm:"start_time"  description:""` //
	FinishTime *gtime.Time `json:"finishTime" orm:"finish_time" description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
