package dto_level

type Edit struct {
	Id         int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"等级id"`
	Icon       string `p:"icon" v:"required#请上传等级图标" dc:"等级图标"`
	Name       string `p:"name" v:"required#请输入等级名称" dc:"等级名称"`
	Level      int    `p:"level" v:"required#请输入等级" dc:"等级"`
	Experience int    `p:"experience" v:"required#请输入成长经验值" dc:"成长经验值"`
	// Benefits   string `p:"benefits" v:"required#请输入等级权益说明" dc:"等级权益说明"`
}
