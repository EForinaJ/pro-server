package dao_user

type Grade struct {
	Id   int64  `json:"id" dc:"等级ID"`
	Name string `json:"name" dc:"等级名称"`
	Icon string `json:"icon" dc:"等级图标"`
}
