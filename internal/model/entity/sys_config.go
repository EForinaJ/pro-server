// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id         int         `json:"id"         orm:"id"          description:"参数主键"` // 参数主键
	Name       string      `json:"name"       orm:"name"        description:"参数名称"` // 参数名称
	Key        string      `json:"key"        orm:"key"         description:"参数键名"` // 参数键名
	Value      string      `json:"value"      orm:"value"       description:"参数键值"` // 参数键值
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"` // 更新时间
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`   // 备注
}
