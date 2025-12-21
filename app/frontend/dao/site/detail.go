package dao_site

type Detail struct {
	Title           string   `json:"title" dc:"网站标题"`
	Domain          string   `json:"domain" dc:"网站域名"`
	Logo            string   `json:"logo" dc:"网站Logo图片地址"`
	Icon            string   `json:"icon" dc:"网站图标地址"`
	Description     string   `json:"description" dc:"网站描述"`
	Address         string   `json:"address" dc:"网站地址"`
	ICP             string   `json:"icp" dc:"ICP备案号"`
	Copyright       string   `json:"copyright" dc:"版权信息"`
	MinFileSize     int      `json:"minFileSize" dc:"文件最小上传体积"`
	BigFileSize     int      `json:"bigFileSize" dc:"文件最大上传体积"`
	MediaType       []string `json:"mediaType" dc:"所支持的文件上传类型"`
	Symbol          string   `json:"symbol" dc:"支付符号，表示支付类型或货币类型"`
	WithDrawPercent int      `json:"withDrawPercent" dc:"提现费率"`
	MinWithdraw     float64  `json:"minWithdraw" dc:"最小提现额度"`
	Contact         *Contact `json:"contact" dc:"联系fangshi"`
}
