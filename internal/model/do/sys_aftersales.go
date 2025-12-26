// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAftersales is the golang structure of table sys_aftersales for DAO operations like Where/Data.
type SysAftersales struct {
	g.Meta     `orm:"table:sys_aftersales, do:true"`
	Id         interface{} //
	Code       interface{} //
	UserId     interface{} //
	ManageId   interface{} //
	OrderId    interface{} //
	Type       interface{} //
	Amount     interface{} //
	Reason     interface{} //
	Status     interface{} //
	CreateTime *gtime.Time //
}
