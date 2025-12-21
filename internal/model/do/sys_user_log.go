// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserLog is the golang structure of table sys_user_log for DAO operations like Where/Data.
type SysUserLog struct {
	g.Meta     `orm:"table:sys_user_log, do:true"`
	Id         interface{} //
	UserId     interface{} //
	ManageId   interface{} //
	Mode       interface{} //
	Content    interface{} //
	CreateTime *gtime.Time //
}
