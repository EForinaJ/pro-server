// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWithdraw is the golang structure of table sys_withdraw for DAO operations like Where/Data.
type SysWithdraw struct {
	g.Meta        `orm:"table:sys_withdraw, do:true"`
	Id            interface{} //
	Code          interface{} //
	ManageId      interface{} //
	UserId        interface{} //
	Money         interface{} //
	SettledAmount interface{} //
	ServiceFee    interface{} //
	ReceiptFiles  interface{} //
	ReceiptNum    interface{} //
	Type          interface{} //
	Number        interface{} //
	Name          interface{} //
	Status        interface{} //
	Remark        interface{} //
	CreateTime    *gtime.Time //
	UpdateTime    *gtime.Time //
}
