// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta     `orm:"table:sys_dict_data, do:true"`
	Id         interface{} //
	Code       interface{} //
	Name       interface{} //
	Key        interface{} //
	Value      interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
