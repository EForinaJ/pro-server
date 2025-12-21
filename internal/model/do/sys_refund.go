// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRefund is the golang structure of table sys_refund for DAO operations like Where/Data.
type SysRefund struct {
	g.Meta     `orm:"table:sys_refund, do:true"`
	Id         interface{} //
	OrderId    interface{} //
	ManageId   interface{} //
	Money      interface{} //
	Mode       interface{} //
	CreateTime *gtime.Time //
}
