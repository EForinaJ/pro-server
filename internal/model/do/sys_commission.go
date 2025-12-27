// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCommission is the golang structure of table sys_commission for DAO operations like Where/Data.
type SysCommission struct {
	g.Meta     `orm:"table:sys_commission, do:true"`
	Id         interface{} //
	WitkeyId   interface{} //
	Related    interface{} //
	After      interface{} //
	Before     interface{} //
	Amount     interface{} //
	Mode       interface{} //
	Type       interface{} //
	Remark     interface{} //
	CreateTime *gtime.Time //
}
