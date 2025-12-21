package dao_site

type Detail struct {
	Title       string   `json:"title" dc:"网站标题"`
	Domain      string   `json:"domain" dc:"网站域名"`
	Logo        string   `json:"logo" dc:"网站Logo图片地址"`
	Icon        string   `json:"icon" dc:"网站图标地址"`
	Description string   `json:"description" dc:"网站描述"`
	Address     string   `json:"address" dc:"网站地址"`
	ICP         string   `json:"icp" dc:"ICP备案号"`
	Copyright   string   `json:"copyright" dc:"版权信息"`
	FileSize    int      `json:"fileSize" dc:"文件最大上传体积"`
	FileType    []string `json:"fileType" dc:"所支持的文件上传类型"`
	ImageSize   int      `json:"imageSize" dc:"图片最大上传体积"`
	ImageType   []string `json:"imageType" dc:"所支持的图片上传类型"`
	Symbol      string   `json:"symbol" dc:"支付符号，表示支付类型或货币类型"`
}
