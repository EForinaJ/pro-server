// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductSpecValues is the golang structure for table sys_product_spec_values.
type SysProductSpecValues struct {
	Id         int64       `json:"id"         orm:"id"           description:""` //
	SpecAttsId int64       `json:"specAttsId" orm:"spec_atts_id" description:""` //
	Value      string      `json:"value"      orm:"value"        description:""` //
	Sort       int         `json:"sort"       orm:"sort"         description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"  description:""` //
}
