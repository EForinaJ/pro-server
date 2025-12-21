// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSku is the golang structure for table sys_product_sku.
type SysProductSku struct {
	Id            int64       `json:"id"            orm:"id"             description:""`             //
	ProductId     int64       `json:"productId"     orm:"product_id"     description:"商品ID"`         // 商品ID
	Code          string      `json:"code"          orm:"code"           description:""`             //
	Price         float64     `json:"price"         orm:"price"          description:"价格"`           // 价格
	OriginalPrice float64     `json:"originalPrice" orm:"original_price" description:""`             //
	Status        int         `json:"status"        orm:"status"         description:"状态：1-禁用，2-启用"` // 状态：1-禁用，2-启用
	Stock         int         `json:"stock"         orm:"stock"          description:""`             //
	CreatedTime   *gtime.Time `json:"createdTime"   orm:"created_time"   description:""`             //
	UpdatedTime   *gtime.Time `json:"updatedTime"   orm:"updated_time"   description:""`             //
}
