package dao_system

type WechatMiniProgram struct {
	Appid  string `json:"appId" dc:"微信小程序AppID"`
	Secret string `json:"secret" dc:"微信小程序Secret"`
}
