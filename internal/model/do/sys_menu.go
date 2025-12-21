// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta        `orm:"table:sys_menu, do:true"`
	Id            interface{} // 菜单ID
	PId           interface{} // 父菜单ID
	MenuType      interface{} // 菜单类型（1目录 2菜单）
	Name          interface{} // 请求地址
	Path          interface{} //
	Label         interface{} // 菜单名称
	Component     interface{} // 组件地址
	Icon          interface{} //
	IsEnable      interface{} //
	Sort          interface{} //
	IsMenu        interface{} //
	KeepAlive     interface{} //
	IsHide        interface{} // 菜单状态（2显示 1隐藏）
	IsHideTab     interface{} //
	Link          interface{} // 跳转
	IsIframe      interface{} //
	ShowBadge     interface{} //
	ShowTextBadge interface{} //
	FixedTab      interface{} //
	ActivePath    interface{} //
	IsFullPage    interface{} //
	AuthName      interface{} //
	AuthLabel     interface{} //
	AuthIcon      interface{} //
	AuthSort      interface{} //
}
