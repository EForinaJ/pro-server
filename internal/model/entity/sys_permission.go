// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure for table sys_permission.
type SysPermission struct {
	Id          int64       `json:"id"          orm:"id"          description:"菜单ID"` // 菜单ID
	Pid         int64       `json:"pid"         orm:"pid"         description:""`     //
	Permission  string      `json:"permission"  orm:"permission"  description:"权限标识"` // 权限标识
	Name        string      `json:"name"        orm:"name"        description:""`     //
	Description string      `json:"description" orm:"description" description:"菜单名称"` // 菜单名称
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新时间"` // 更新时间
}
