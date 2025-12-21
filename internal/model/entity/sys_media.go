// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMedia is the golang structure for table sys_media.
type SysMedia struct {
	Id         int64       `json:"id"         orm:"id"          description:""`       //
	Path       string      `json:"path"       orm:"path"        description:"存放路径"`   // 存放路径
	Name       string      `json:"name"       orm:"name"        description:"文件名称"`   // 文件名称
	OrName     string      `json:"orName"     orm:"or_name"     description:"原始文件名称"` // 原始文件名称
	Size       string      `json:"size"       orm:"size"        description:"文件大小"`   // 文件大小
	Ext        string      `json:"ext"        orm:"ext"         description:"文件后缀"`   // 文件后缀
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`   // 创建时间
	DeleteTime *gtime.Time `json:"deleteTime" orm:"delete_time" description:""`       //
	MediaType  string      `json:"mediaType"  orm:"media_type"  description:"文件类型"`   // 文件类型
}
