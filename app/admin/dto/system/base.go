package dto_system

type Base struct {
	Title       string `v:"required#请输入站点标题" json:"title" p:"title" dc:"网站标题"`
	Domain      string `v:"required#请输入站点域名" json:"domain" p:"domain" dc:"网站域名"`
	Logo        string `v:"required#请上传站点Logo地址" json:"logo" p:"logo" dc:"网站Logo图片地址"`
	Icon        string `v:"required#请上传站点图标地址" json:"icon" p:"icon" dc:"网站图标地址"`
	Description string `v:"required#请输入站点描述" json:"description" p:"description" dc:"网站描述"`
	ICP         string `v:"required#请输入ICP备案号" json:"icp" p:"icp" dc:"ICP备案号"`
}
