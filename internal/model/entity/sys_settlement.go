// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSettlement is the golang structure for table sys_settlement.
type SysSettlement struct {
	Id             int64       `json:"id"             orm:"id"              description:""` //
	OrderId        int64       `json:"orderId"        orm:"order_id"        description:""` //
	WitkeyId       int64       `json:"witkeyId"       orm:"witkey_id"       description:""` //
	ManageId       int64       `json:"manageId"       orm:"manage_id"       description:""` //
	Type           int         `json:"type"           orm:"type"            description:""` //
	Amount         float64     `json:"amount"         orm:"amount"          description:""` //
	Commission     float64     `json:"commission"     orm:"commission"      description:""` //
	Deduction      float64     `json:"deduction"      orm:"deduction"       description:""` //
	ServiceCharge  float64     `json:"serviceCharge"  orm:"service_charge"  description:""` //
	Status         int         `json:"status"         orm:"status"          description:""` //
	Remark         string      `json:"remark"         orm:"remark"          description:""` //
	SettlementTime *gtime.Time `json:"settlementTime" orm:"settlement_time" description:""` //
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:""` //
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"     description:""` //
}
