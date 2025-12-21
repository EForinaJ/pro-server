// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserLoginLog is the golang structure of table sys_user_login_log for DAO operations like Where/Data.
type SysUserLoginLog struct {
	g.Meta     `orm:"table:sys_user_login_log, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Ip         interface{} //
	CreateTime *gtime.Time //
}
