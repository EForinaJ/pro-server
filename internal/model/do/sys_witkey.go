// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkey is the golang structure of table sys_witkey for DAO operations like Where/Data.
type SysWitkey struct {
	g.Meta      `orm:"table:sys_witkey, do:true"`
	Id          interface{} // 用户ID
	UserId      interface{} //
	TitleId     interface{} //
	GameId      interface{} //
	Album       interface{} //
	Rate        interface{} //
	Commission  interface{} //
	Description interface{} //
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
