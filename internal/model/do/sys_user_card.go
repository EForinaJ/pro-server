// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserCard is the golang structure of table sys_user_card for DAO operations like Where/Data.
type SysUserCard struct {
	g.Meta     `orm:"table:sys_user_card, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Type       interface{} //
	Name       interface{} //
	Number     interface{} //
	NickName   interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
