package dto_media

type Query struct {
	Page      int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit     int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Drive     string `p:"drive" dc:"上传的驱动"`
	MediaType string `p:"mediaType" v:"in:image,video,audio,file#附件类型错误" dc:"附件类型"`
	Name      string `p:"name" dc:"附件名称"`
}
