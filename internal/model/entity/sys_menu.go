// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id            int64  `json:"id"            orm:"id"              description:"菜单ID"`          // 菜单ID
	PId           int64  `json:"pId"           orm:"p_id"            description:"父菜单ID"`         // 父菜单ID
	MenuType      string `json:"menuType"      orm:"menu_type"       description:"菜单类型（1目录 2菜单）"` // 菜单类型（1目录 2菜单）
	Name          string `json:"name"          orm:"name"            description:"请求地址"`          // 请求地址
	Path          string `json:"path"          orm:"path"            description:""`              //
	Label         string `json:"label"         orm:"label"           description:"菜单名称"`          // 菜单名称
	Component     string `json:"component"     orm:"component"       description:"组件地址"`          // 组件地址
	Icon          string `json:"icon"          orm:"icon"            description:""`              //
	IsEnable      int    `json:"isEnable"      orm:"is_enable"       description:""`              //
	Sort          int    `json:"sort"          orm:"sort"            description:""`              //
	IsMenu        int    `json:"isMenu"        orm:"is_menu"         description:""`              //
	KeepAlive     int    `json:"keepAlive"     orm:"keep_alive"      description:""`              //
	IsHide        int    `json:"isHide"        orm:"is_hide"         description:"菜单状态（2显示 1隐藏）"` // 菜单状态（2显示 1隐藏）
	IsHideTab     int    `json:"isHideTab"     orm:"is_hide_tab"     description:""`              //
	Link          string `json:"link"          orm:"link"            description:"跳转"`            // 跳转
	IsIframe      int    `json:"isIframe"      orm:"is_iframe"       description:""`              //
	ShowBadge     int    `json:"showBadge"     orm:"show_badge"      description:""`              //
	ShowTextBadge int    `json:"showTextBadge" orm:"show_text_badge" description:""`              //
	FixedTab      int    `json:"fixedTab"      orm:"fixed_tab"       description:""`              //
	ActivePath    string `json:"activePath"    orm:"active_path"     description:""`              //
	IsFullPage    int    `json:"isFullPage"    orm:"is_full_page"    description:""`              //
	AuthName      string `json:"authName"      orm:"auth_name"       description:""`              //
	AuthLabel     string `json:"authLabel"     orm:"auth_label"      description:""`              //
	AuthIcon      string `json:"authIcon"      orm:"auth_icon"       description:""`              //
	AuthSort      int    `json:"authSort"      orm:"auth_sort"       description:""`              //
}
