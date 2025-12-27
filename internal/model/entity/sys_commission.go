// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCommission is the golang structure for table sys_commission.
type SysCommission struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	WitkeyId   int64       `json:"witkeyId"   orm:"witkey_id"   description:""` //
	Related    string      `json:"related"    orm:"related"     description:""` //
	After      float64     `json:"after"      orm:"after"       description:""` //
	Before     float64     `json:"before"     orm:"before"      description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	Mode       int         `json:"mode"       orm:"mode"        description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Remark     string      `json:"remark"     orm:"remark"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
