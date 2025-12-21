// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPayment is the golang structure of table sys_payment for DAO operations like Where/Data.
type SysPayment struct {
	g.Meta     `orm:"table:sys_payment, do:true"`
	Id         interface{} //
	RelatedId  interface{} //
	Code       interface{} //
	Type       interface{} //
	Money      interface{} //
	Mode       interface{} //
	CreateTime *gtime.Time //
}
