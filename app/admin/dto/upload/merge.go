package dto_upload

type Merge struct {
	Identifier string `p:"identifier"  v:"required#请上传文件唯一标识" dc:"切片标识"`
	FileName   string `p:"fileName"  v:"required#请设置文件名" dc:"文件名"`
	Size       int64  `p:"size"  v:"required#请设置文件大小" dc:"文件大小"`
}
