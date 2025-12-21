// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProject is the golang structure of table sys_project for DAO operations like Where/Data.
type SysProject struct {
	g.Meta     `orm:"table:sys_project, do:true"`
	Id         interface{} //
	Code       interface{} //
	OrderId    interface{} //
	WitkeyId   interface{} //
	Images     interface{} //
	Money      interface{} //
	ServiceFee interface{} //
	Commission interface{} //
	Status     interface{} //
	StartTime  *gtime.Time //
	FinishTime *gtime.Time //
	CreateTime *gtime.Time //
}
