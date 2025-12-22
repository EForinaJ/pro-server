// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderWitkey is the golang structure of table sys_order_witkey for DAO operations like Where/Data.
type SysOrderWitkey struct {
	g.Meta     `orm:"table:sys_order_witkey, do:true"`
	OrderId    interface{} //
	WitkeyId   interface{} //
	IsReplaced interface{} //
	Reason     interface{} //
	CreateTime *gtime.Time //
}
