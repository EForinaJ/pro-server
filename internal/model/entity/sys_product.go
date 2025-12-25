// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProduct is the golang structure for table sys_product.
type SysProduct struct {
	Id            int64       `json:"id"            orm:"id"             description:""`             //
	Pic           string      `json:"pic"           orm:"pic"            description:""`             //
	Type          int         `json:"type"          orm:"type"           description:"1护航，2代肝"`      // 1护航，2代肝
	CategoryId    int64       `json:"categoryId"    orm:"category_id"    description:""`             //
	GameId        int64       `json:"gameId"        orm:"game_id"        description:""`             //
	Name          string      `json:"name"          orm:"name"           description:""`             //
	Description   string      `json:"description"   orm:"description"    description:""`             //
	Content       string      `json:"content"       orm:"content"        description:""`             //
	WitkeyCount   int         `json:"witkeyCount"   orm:"witkey_count"   description:""`             //
	Price         float64     `json:"price"         orm:"price"          description:""`             //
	Unit          string      `json:"unit"          orm:"unit"           description:""`             //
	SalesCount    int64       `json:"salesCount"    orm:"sales_count"    description:""`             //
	Rate          float64     `json:"rate"          orm:"rate"           description:""`             //
	PurchaseLimit int         `json:"purchaseLimit" orm:"purchase_limit" description:""`             //
	SkuType       int         `json:"skuType"       orm:"sku_type"       description:""`             //
	Status        int         `json:"status"        orm:"status"         description:"状态：1-禁用，2-启用"` // 状态：1-禁用，2-启用
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    description:""`             //
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"    description:""`             //
}
