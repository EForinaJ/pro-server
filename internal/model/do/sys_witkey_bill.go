// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyBill is the golang structure of table sys_witkey_bill for DAO operations like Where/Data.
type SysWitkeyBill struct {
	g.Meta     `orm:"table:sys_witkey_bill, do:true"`
	Id         interface{} //
	WitkeyId   interface{} //
	RelatedId  interface{} //
	Code       interface{} //
	Type       interface{} //
	Money      interface{} //
	Mode       interface{} //
	CreateTime *gtime.Time //
}
