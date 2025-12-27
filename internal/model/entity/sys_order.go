// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrder is the golang structure for table sys_order.
type SysOrder struct {
	Id             int64       `json:"id"             orm:"id"              description:""`     //
	Code           string      `json:"code"           orm:"code"            description:""`     //
	UserId         int64       `json:"userId"         orm:"user_id"         description:""`     //
	ProductId      int64       `json:"productId"      orm:"product_id"      description:""`     //
	ManageId       int64       `json:"manageId"       orm:"manage_id"       description:""`     //
	Specifications string      `json:"specifications" orm:"specifications"  description:""`     //
	WitkeyCount    int         `json:"witkeyCount"    orm:"witkey_count"    description:""`     //
	Quantity       int         `json:"quantity"       orm:"quantity"        description:""`     //
	Unit           string      `json:"unit"           orm:"unit"            description:""`     //
	UnitPrice      float64     `json:"unitPrice"      orm:"unit_price"      description:""`     //
	TotalAmount    float64     `json:"totalAmount"    orm:"total_amount"    description:"订单总额"` // 订单总额
	DiscountAmount float64     `json:"discountAmount" orm:"discount_amount" description:"优惠金额"` // 优惠金额
	ActualAmount   float64     `json:"actualAmount"   orm:"actual_amount"   description:"实付金额"` // 实付金额
	PayAmount      float64     `json:"payAmount"      orm:"pay_amount"      description:""`     //
	Requirements   string      `json:"requirements"   orm:"requirements"    description:""`     //
	Status         int         `json:"status"         orm:"status"          description:""`     //
	PayCode        string      `json:"payCode"        orm:"pay_code"        description:""`     //
	PayMode        int         `json:"payMode"        orm:"pay_mode"        description:""`     //
	PayStatus      int         `json:"payStatus"      orm:"pay_status"      description:""`     //
	PayTime        *gtime.Time `json:"payTime"        orm:"pay_time"        description:""`     //
	StartTime      *gtime.Time `json:"startTime"      orm:"start_time"      description:""`     //
	FinishTime     *gtime.Time `json:"finishTime"     orm:"finish_time"     description:""`     //
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:""`     //
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"     description:""`     //
	Remark         string      `json:"remark"         orm:"remark"          description:""`     //
}
