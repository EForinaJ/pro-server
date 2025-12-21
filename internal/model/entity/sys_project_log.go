// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProjectLog is the golang structure for table sys_project_log.
type SysProjectLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	ProjectId  int64       `json:"projectId"  orm:"project_id"  description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
