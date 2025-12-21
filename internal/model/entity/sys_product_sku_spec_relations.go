// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSkuSpecRelations is the golang structure for table sys_product_sku_spec_relations.
type SysProductSkuSpecRelations struct {
	Id          int64       `json:"id"          orm:"id"            description:""` //
	SkuId       int64       `json:"skuId"       orm:"sku_id"        description:""` //
	SpecAttrId  int64       `json:"specAttrId"  orm:"spec_attr_id"  description:""` //
	SpecValueId int64       `json:"specValueId" orm:"spec_value_id" description:""` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:""` //
}
