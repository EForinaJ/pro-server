// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMedia is the golang structure of table sys_media for DAO operations like Where/Data.
type SysMedia struct {
	g.Meta     `orm:"table:sys_media, do:true"`
	Id         interface{} //
	Path       interface{} // 存放路径
	Name       interface{} // 文件名称
	OrName     interface{} // 原始文件名称
	Size       interface{} // 文件大小
	Ext        interface{} // 文件后缀
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time //
	MediaType  interface{} // 文件类型
}
