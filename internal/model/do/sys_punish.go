// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPunish is the golang structure of table sys_punish for DAO operations like Where/Data.
type SysPunish struct {
	g.Meta     `orm:"table:sys_punish, do:true"`
	Id         interface{} //
	Code       interface{} //
	WitkeyId   interface{} //
	ManageId   interface{} //
	Money      interface{} //
	Content    interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
