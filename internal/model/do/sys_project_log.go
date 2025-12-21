// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProjectLog is the golang structure of table sys_project_log for DAO operations like Where/Data.
type SysProjectLog struct {
	g.Meta     `orm:"table:sys_project_log, do:true"`
	Id         interface{} //
	ManageId   interface{} //
	ProjectId  interface{} //
	Type       interface{} //
	Content    interface{} //
	CreateTime *gtime.Time //
}
