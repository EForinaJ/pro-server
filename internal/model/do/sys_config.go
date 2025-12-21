// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta     `orm:"table:sys_config, do:true"`
	Id         interface{} // 参数主键
	Name       interface{} // 参数名称
	Key        interface{} // 参数键名
	Value      interface{} // 参数键值
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Remark     interface{} // 备注
}
