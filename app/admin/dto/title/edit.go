package dto_title

type Edit struct {
	Id             int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"头衔id"`
	Pic            string `p:"pic" v:"required#请上传分类图片" dc:"分类图片"`
	Name           string `p:"name" v:"required#请输入分类名称" dc:"分类名称"`
	Description    string `p:"description" dc:"分类描述"`
	ServicePercent int    `p:"servicePercent" v:"required#请输入服务百分比" dc:"服务百分比"`
	GameId         int64  `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属有游戏id"`
}
