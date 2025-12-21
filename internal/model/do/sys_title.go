// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTitle is the golang structure of table sys_title for DAO operations like Where/Data.
type SysTitle struct {
	g.Meta         `orm:"table:sys_title, do:true"`
	Id             interface{} //
	GameId         interface{} //
	Name           interface{} //
	ServicePercent interface{} //
	Pic            interface{} //
	Description    interface{} //
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
}
