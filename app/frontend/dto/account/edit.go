package dto_account

type Edit struct {
	Name        string `p:"name" v:"required#请输入修改名称" dc:"名称"`
	Sex         int    `p:"sex"     dc:"性别"`
	Address     string `p:"address" dc:"所在地"`
	Birthday    int64  `p:"birthday" dc:"生日"`
	Avatar      string `p:"avatar" dc:"头像地址"`
	Description string `p:"description" dc:"简介"`
}
