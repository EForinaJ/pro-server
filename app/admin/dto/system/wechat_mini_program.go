package dto_system

type WechatMiniProgram struct {
	Appid  string `v:"required#请输入微信小程序AppID" json:"appId" p:"appId" dc:"微信小程序AppID"`
	Secret string `v:"required#请输入微信小程序AppSecret" json:"secret" p:"secret" dc:"微信小程序secret"`
}
