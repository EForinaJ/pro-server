package dao_system

type Base struct {
	Title       string `json:"title" dc:"网站标题"`
	Domain      string `json:"domain" dc:"网站域名"`
	Logo        string `json:"logo" dc:"网站Logo图片地址"`
	Icon        string `json:"icon" dc:"网站图标地址"`
	Description string `json:"description" dc:"网站描述"`
	Address     string `json:"address" dc:"网站地址"`
	ICP         string `json:"icp" dc:"ICP备案号"`
	Copyright   string `json:"copyright" dc:"版权信息"`
}
