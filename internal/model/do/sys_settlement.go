// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSettlement is the golang structure of table sys_settlement for DAO operations like Where/Data.
type SysSettlement struct {
	g.Meta         `orm:"table:sys_settlement, do:true"`
	Id             interface{} //
	Code           interface{} //
	OrderId        interface{} //
	WitkeyId       interface{} //
	ManageId       interface{} //
	Amount         interface{} //
	Commission     interface{} //
	ServiceCharge  interface{} //
	Status         interface{} //
	Reason         interface{} //
	SettlementTime *gtime.Time //
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
}
