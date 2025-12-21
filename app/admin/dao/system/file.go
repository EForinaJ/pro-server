package dao_system

type File struct {
	Engine    string   `json:"engine" dc:"上传平台"`
	Path      string   `json:"path" dc:"上传路径"`
	FileSize  int      `json:"fileSize" dc:"文件最大上传体积"`
	FileType  []string `json:"fileType" dc:"所支持的文件上传类型"`
	ImageSize int      `json:"imageSize" dc:"图片最大上传体积"`
	ImageType []string `json:"imageType" dc:"所支持的图片上传类型"`
}
