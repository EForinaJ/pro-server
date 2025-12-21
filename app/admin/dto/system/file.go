package dto_system

type File struct {
	Engine    string   `json:"engine" p:"engine" v:"required#上传平台不能为空" dc:"上传平台"`
	Path      string   `json:"path" p:"path" v:"required#上传路径不能为空" dc:"上传路径"`
	FileSize  int      `json:"fileSize" p:"fileSize" v:"required#文件最大上传体积不能为空" dc:"文件最大上传体积"`
	FileType  []string `json:"fileType" p:"fileType" v:"required|array#所支持的文件上传类型不能为空|所支持的文件上传类型必须是一个字符串列表" dc:"所支持的文件上传类型"`
	ImageSize int      `json:"imageSize" p:"imageSize" v:"required#图片最大上传体积不能为空" dc:"图片最大上传体积"`
	ImageType []string `json:"imageType" p:"imageType" v:"required|array#所支持的图片上传类型不能为空|所支持的图片上传类型必须是一个字符串列表" dc:"所支持的图片上传类型"`
}
