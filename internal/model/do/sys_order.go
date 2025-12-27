// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrder is the golang structure of table sys_order for DAO operations like Where/Data.
type SysOrder struct {
	g.Meta         `orm:"table:sys_order, do:true"`
	Id             interface{} //
	Code           interface{} //
	UserId         interface{} //
	ProductId      interface{} //
	ManageId       interface{} //
	Specifications interface{} //
	WitkeyCount    interface{} //
	Quantity       interface{} //
	Unit           interface{} //
	UnitPrice      interface{} //
	TotalAmount    interface{} // 订单总额
	DiscountAmount interface{} // 优惠金额
	ActualAmount   interface{} // 实付金额
	PayAmount      interface{} //
	Requirements   interface{} //
	Status         interface{} //
	PayCode        interface{} //
	PayMode        interface{} //
	PayStatus      interface{} //
	PayTime        *gtime.Time //
	StartTime      *gtime.Time //
	FinishTime     *gtime.Time //
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
	Remark         interface{} //
}
