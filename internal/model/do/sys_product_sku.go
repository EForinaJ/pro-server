// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSku is the golang structure of table sys_product_sku for DAO operations like Where/Data.
type SysProductSku struct {
	g.Meta        `orm:"table:sys_product_sku, do:true"`
	Id            interface{} //
	ProductId     interface{} // 商品ID
	Code          interface{} //
	Price         interface{} // 价格
	OriginalPrice interface{} //
	Status        interface{} // 状态：1-禁用，2-启用
	Stock         interface{} //
	CreatedTime   *gtime.Time //
	UpdatedTime   *gtime.Time //
}
