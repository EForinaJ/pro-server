// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRecharge is the golang structure of table sys_recharge for DAO operations like Where/Data.
type SysRecharge struct {
	g.Meta     `orm:"table:sys_recharge, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Code       interface{} //
	Amount     interface{} //
	PayMode    interface{} //
	Status     interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	Remark     interface{} //
}
