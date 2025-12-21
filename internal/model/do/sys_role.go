// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta      `orm:"table:sys_role, do:true"`
	Id          interface{} // 角色ID
	Name        interface{} // 角色名称
	Code        interface{} //
	Description interface{} //
	Status      interface{} // 角色状态（2正常 1停用）
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	Remark      interface{} // 备注
}
