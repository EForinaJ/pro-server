package dto_user

type Create struct {
	Phone string `p:"phone" v:"required|phone#请设置手机号码|手机号码格式错误"     dc:"手机号码"`
	// LevelId     int64  `p:"LevelId"  dc:"用户等级"`
	Name        string   `p:"name" dc:"名称"`
	Avatar      string   `p:"avatar"  dc:"头像"`
	Cover       string   `p:"cover"  dc:"封面"`
	Address     []string `p:"address"  dc:"地址"`
	Description string   `p:"description"  dc:"描述"`
	Sex         int      `p:"sex"  dc:"性别"`
	Birthday    int64    `p:"birthday" dc:"生日"`
	Password    string   `p:"password"  dc:"密码"`
	Status      int      `p:"status"  dc:"状态"`
}
