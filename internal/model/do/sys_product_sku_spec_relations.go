// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSkuSpecRelations is the golang structure of table sys_product_sku_spec_relations for DAO operations like Where/Data.
type SysProductSkuSpecRelations struct {
	g.Meta      `orm:"table:sys_product_sku_spec_relations, do:true"`
	Id          interface{} //
	SkuId       interface{} //
	SpecAttrId  interface{} //
	SpecValueId interface{} //
	CreateTime  *gtime.Time //
}
