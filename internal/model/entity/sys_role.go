// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id          int64       `json:"id"          orm:"id"          description:"角色ID"`          // 角色ID
	Name        string      `json:"name"        orm:"name"        description:"角色名称"`          // 角色名称
	Code        string      `json:"code"        orm:"code"        description:""`              //
	Description string      `json:"description" orm:"description" description:""`              //
	Status      int         `json:"status"      orm:"status"      description:"角色状态（2正常 1停用）"` // 角色状态（2正常 1停用）
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"`          // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新时间"`          // 更新时间
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`            // 备注
}
