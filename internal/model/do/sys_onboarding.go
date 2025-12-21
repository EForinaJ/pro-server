// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOnboarding is the golang structure of table sys_onboarding for DAO operations like Where/Data.
type SysOnboarding struct {
	g.Meta       `orm:"table:sys_onboarding, do:true"`
	Id           interface{} //
	ManageId     interface{} //
	UserId       interface{} //
	TitleId      interface{} //
	GameId       interface{} //
	Images       interface{} //
	WxNumber     interface{} //
	Status       interface{} //
	RefuseReason interface{} //
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
}
