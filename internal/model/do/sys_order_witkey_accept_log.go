// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderWitkeyAcceptLog is the golang structure of table sys_order_witkey_accept_log for DAO operations like Where/Data.
type SysOrderWitkeyAcceptLog struct {
	g.Meta     `orm:"table:sys_order_witkey_accept_log, do:true"`
	OrderId    interface{} //
	WitkeyId   interface{} //
	CreateTime *gtime.Time //
}
