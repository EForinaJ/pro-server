// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure of table sys_comment for DAO operations like Where/Data.
type SysComment struct {
	g.Meta      `orm:"table:sys_comment, do:true"`
	Id          interface{} //
	UserId      interface{} //
	OrderId     interface{} //
	ProductId   interface{} //
	ParentId    interface{} //
	ReplyUserId interface{} //
	Content     interface{} //
	Rating      interface{} //
	Images      interface{} //
	IsTop       interface{} //
	Status      interface{} //
	Ip          interface{} //
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
}
