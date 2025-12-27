// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCapital is the golang structure of table sys_capital for DAO operations like Where/Data.
type SysCapital struct {
	g.Meta     `orm:"table:sys_capital, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Related    interface{} //
	Code       interface{} //
	Type       interface{} //
	Amount     interface{} //
	Mode       interface{} //
	CreateTime *gtime.Time //
}
