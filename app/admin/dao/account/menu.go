package dao_account

type Menu struct {
	Id            int64  `json:"id" description:"菜单ID"`
	PId           int64  `json:"pId" description:"父菜单ID"`
	MenuType      string `json:"menuType" description:"菜单类型"`
	Name          string `json:"name" description:"请求地址"`
	Path          string `json:"path" description:"路由地址"`
	Label         string `json:"label" description:"菜单名称"`
	Component     string `json:"component" description:"组件地址"`
	Icon          string `json:"icon" description:"菜单图标"`
	IsEnable      bool   `json:"isEnable" description:"是否启用"`
	Sort          int    `json:"sort" description:"排序"`
	IsMenu        bool   `json:"isMenu" description:"是否为目录"`
	KeepAlive     bool   `json:"keepAlive" description:"缓存"`
	IsHide        bool   `json:"isHide" description:"是否显示菜单"`
	IsHideTab     bool   `json:"isHideTab" description:"是否显示标签栏"`
	Link          string `json:"link" description:"跳转"`
	IsIframe      bool   `json:"isIframe" description:"是否Iframe"`
	ShowBadge     bool   `json:"showBadge" description:"显示图标"`
	ShowTextBadge bool   `json:"showTextBadge" description:"显示文字图标"`
	FixedTab      bool   `json:"fixedTab" description:"固定标签栏"`
	ActivePath    string `json:"activePath" description:"所激活的路由"`
	IsFullPage    bool   `json:"isFullPage" description:"是否全屏"`
	AuthName      string `json:"authName" description:"权限名称"`
	AuthLabel     string `json:"authLabel" description:"权限标识"`
	AuthIcon      string `json:"authIcon" description:"权限图标"`
	AuthSort      int    `json:"authSort" description:"权限排序"`
}
