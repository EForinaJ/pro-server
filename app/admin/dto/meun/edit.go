package dto_meun

type Edit struct {
	Id            int64  `p:"id" v:"required|integer|min:1#请输入菜单id|id类型必须是整型|id最小为1" dc:"菜单id"`
	PId           int64  `p:"pId" description:"父菜单ID"`
	MenuType      string `p:"menuType" v:"required#请设置菜单类型" description:"菜单类型"`
	Name          string `p:"name" v:"required-if:menuType,menu#请输入菜单名称" description:"菜单名称"`
	Path          string `p:"path" v:"required-if:menuType,menu#请输入路由地址" description:"路由地址"`
	Label         string `p:"label" v:"required-if:menuType,menu#请输入权限标识" description:"权限标识"`
	Component     string `p:"component" description:"组件地址"`
	Icon          string `p:"icon" description:"菜单图标"`
	IsEnable      bool   `p:"isEnable" description:"是否启用"`
	Sort          int    `p:"sort" description:"排序"`
	IsMenu        bool   `p:"isMenu" description:"是否为目录"`
	KeepAlive     bool   `p:"keepAlive" description:"缓存"`
	IsHide        bool   `p:"isHide" description:"是否显示菜单"`
	IsHideTab     bool   `p:"isHideTab" description:"是否显示标签栏"`
	Link          string `p:"link" description:"跳转"`
	IsIframe      bool   `p:"isIframe" description:"是否Iframe"`
	ShowBadge     bool   `p:"showBadge" description:"显示图标"`
	ShowTextBadge bool   `p:"showTextBadge" description:"显示文字图标"`
	FixedTab      bool   `p:"fixedTab" description:"固定标签栏"`
	ActivePath    string `p:"activePath" description:"所激活的路由"`
	IsFullPage    bool   `p:"isFullPage" description:"是否全屏"`
	AuthName      string `p:"authName" v:"required-if:menuType,button#请输入权限名称" description:"权限名称"`
	AuthLabel     string `p:"authLabel" v:"required-if:menuType,button#请输入权限标识" description:"权限标识"`
	AuthIcon      string `p:"authIcon" description:"权限图标"`
	AuthSort      int    `p:"authSort" description:"权限排序"`
}
