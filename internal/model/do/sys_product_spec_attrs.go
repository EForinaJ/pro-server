// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSpecAttrs is the golang structure of table sys_product_spec_attrs for DAO operations like Where/Data.
type SysProductSpecAttrs struct {
	g.Meta     `orm:"table:sys_product_spec_attrs, do:true"`
	Id         interface{} //
	ProductId  interface{} //
	Name       interface{} //
	CreateTime *gtime.Time //
}
