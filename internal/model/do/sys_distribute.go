// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDistribute is the golang structure of table sys_distribute for DAO operations like Where/Data.
type SysDistribute struct {
	g.Meta     `orm:"table:sys_distribute, do:true"`
	Id         interface{} //
	OrderId    interface{} //
	WitkeyId   interface{} //
	ManageId   interface{} //
	IsCancel   interface{} //
	Reason     interface{} //
	CreateTime *gtime.Time //
}
