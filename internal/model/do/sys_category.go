// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCategory is the golang structure of table sys_category for DAO operations like Where/Data.
type SysCategory struct {
	g.Meta      `orm:"table:sys_category, do:true"`
	Id          interface{} //
	GameId      interface{} //
	Pid         interface{} // 顶级分类
	Name        interface{} // 分类名称
	Pic         interface{} // 分类背景图
	Sort        interface{} // 分类排序
	Description interface{} // 分类描述
	CreateTime  *gtime.Time // 创建日期
	UpdateTime  *gtime.Time // 更新日期
}
