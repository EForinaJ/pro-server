package dao_system

type File struct {
	Engine      string   `json:"engine" dc:"上传平台"`
	Path        string   `json:"path" dc:"上传路径"`
	MinFileSize int      `json:"minFileSize" dc:"文件最小上传体积"`
	BigFileSize int      `json:"bigFileSize" dc:"文件最大上传体积"`
	MediaType   []string `json:"mediaType" dc:"所支持的文件上传类型"`
}
