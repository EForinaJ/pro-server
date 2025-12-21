// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSpecValues is the golang structure of table sys_product_spec_values for DAO operations like Where/Data.
type SysProductSpecValues struct {
	g.Meta     `orm:"table:sys_product_spec_values, do:true"`
	Id         interface{} //
	SpecAttsId interface{} //
	Value      interface{} //
	Sort       interface{} //
	CreateTime *gtime.Time //
}
