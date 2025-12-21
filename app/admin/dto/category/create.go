package dto_category

type Create struct {
	Pic         string `p:"pic" v:"required#请上传分类图片" dc:"分类图片"`
	Name        string `p:"name" v:"required#请输入分类名称" dc:"分类名称"`
	Description string `p:"description" dc:"分类描述"`
	GameId      int64  `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属分类id"`
}
