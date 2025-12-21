// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysManageLoginLog is the golang structure for table sys_manage_login_log.
type SysManageLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	Ip         string      `json:"ip"         orm:"ip"          description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
