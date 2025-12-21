// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSpecAttrs is the golang structure for table sys_product_spec_attrs.
type SysProductSpecAttrs struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	ProductId  int64       `json:"productId"  orm:"product_id"  description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
