// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure of table sys_permission for DAO operations like Where/Data.
type SysPermission struct {
	g.Meta      `orm:"table:sys_permission, do:true"`
	Id          interface{} // 菜单ID
	Pid         interface{} //
	Permission  interface{} // 权限标识
	Name        interface{} //
	Description interface{} // 菜单名称
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
