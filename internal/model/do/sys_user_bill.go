// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserBill is the golang structure of table sys_user_bill for DAO operations like Where/Data.
type SysUserBill struct {
	g.Meta     `orm:"table:sys_user_bill, do:true"`
	Id         interface{} //
	UserId     interface{} //
	RelatedId  interface{} //
	Code       interface{} //
	Type       interface{} //
	Amount     interface{} //
	Mode       interface{} //
	CreateTime *gtime.Time //
}
