// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProduct is the golang structure of table sys_product for DAO operations like Where/Data.
type SysProduct struct {
	g.Meta      `orm:"table:sys_product, do:true"`
	Id          interface{} //
	Pic         interface{} //
	Type        interface{} // 1护航，2代肝
	CategoryId  interface{} //
	GameId      interface{} //
	Name        interface{} //
	Description interface{} //
	Content     interface{} //
	WitkeyCount interface{} //
	Price       interface{} //
	Unit        interface{} //
	SalesCount  interface{} //
	Rate        interface{} //
	SkuType     interface{} //
	Status      interface{} // 状态：1-禁用，2-启用
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
}
