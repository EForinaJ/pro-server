// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta     `orm:"table:sys_dict_type, do:true"`
	Id         interface{} //
	Name       interface{} //
	Code       interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
