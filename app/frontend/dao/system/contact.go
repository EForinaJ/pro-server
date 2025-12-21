package dao_system

type Contact struct {
	Platform string `json:"platform" dc:"联系平台"`
	Number   string `json:"number" dc:"平台号吗"`
	Icon     string `json:"icon" dc:"平台图标"`
	Role     string `json:"role" dc:"客服角色标识"`
}
