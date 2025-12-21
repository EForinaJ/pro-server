package dto_game

type Create struct {
	Pic         string `p:"pic" v:"required#请上传分类图片" dc:"分类图片"`
	Name        string `p:"name" v:"required#请输入分类名称" dc:"分类名称"`
	Description string `p:"description" dc:"分类描述"`
}
