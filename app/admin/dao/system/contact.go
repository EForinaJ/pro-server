package dao_system

type Contact struct {
	Platform   string `json:"platform" dc:"平台名称"`
	Number     string `json:"number" dc:"平台号吗"`
	Icon       string `json:"icon" dc:"平台图标"`
	OnlineTime string `json:"onlineTime" dc:"在线时间"`
	Wechat     string `json:"wechat" dc:"微信二维码"`
}
