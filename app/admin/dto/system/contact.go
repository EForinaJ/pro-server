package dto_system

type Contact struct {
	Platform   string `v:"required#请输入联系平台" p:"platform" json:"platform" dc:"平台名称"`
	Number     string `v:"required#请输入平台号吗" p:"number" json:"number" dc:"平台号吗"`
	Icon       string `v:"required#请上传平台图标" p:"icon" json:"icon" dc:"平台图标"`
	OnlineTime string `v:"required#请输在线时间" p:"onlineTime" json:"onlineTime" dc:"在线时间"`
	Wechat     string `v:"required#请输上传平台图标" p:"wechat" json:"wechat" dc:"微信二维码"`
}
