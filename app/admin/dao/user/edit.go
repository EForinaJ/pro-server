package dao_user

type Edit struct {
	Id string `json:"id" dc:"ID"`
	// LevelId     int64    `json:"levelId" dc:"等级"`
	Name     string   `json:"name" dc:"名称"`
	Phone    string   `json:"phone" dc:"手机号码"`
	Sex      int      `json:"sex" dc:"性别"`
	Address  []string `json:"address" dc:"地址"`
	Birthday int64    `json:"birthday" dc:"生日"`
	Avatar   string   `json:"avatar" dc:"头像地址"`
	// Cover       string   `json:"cover" dc:"封面地址"`
	Description string `json:"description" dc:"简介"`
	Status      int    `json:"status" dc:"状态"`
}
