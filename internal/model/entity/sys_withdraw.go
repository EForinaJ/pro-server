// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWithdraw is the golang structure for table sys_withdraw.
type SysWithdraw struct {
	Id            int64       `json:"id"            orm:"id"             description:""` //
	Code          string      `json:"code"          orm:"code"           description:""` //
	ManageId      int64       `json:"manageId"      orm:"manage_id"      description:""` //
	UserId        int64       `json:"userId"        orm:"user_id"        description:""` //
	Money         float64     `json:"money"         orm:"money"          description:""` //
	SettledAmount float64     `json:"settledAmount" orm:"settled_amount" description:""` //
	ServiceFee    float64     `json:"serviceFee"    orm:"service_fee"    description:""` //
	ReceiptFiles  string      `json:"receiptFiles"  orm:"receipt_files"  description:""` //
	ReceiptNum    string      `json:"receiptNum"    orm:"receipt_num"    description:""` //
	Type          int         `json:"type"          orm:"type"           description:""` //
	Number        string      `json:"number"        orm:"number"         description:""` //
	Name          string      `json:"name"          orm:"name"           description:""` //
	Status        int         `json:"status"        orm:"status"         description:""` //
	Remark        string      `json:"remark"        orm:"remark"         description:""` //
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:""` //
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"    description:""` //
}
