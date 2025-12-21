package dao_level

type List struct {
	Id         int64  `json:"id" dc:"ID"`
	Name       string `json:"name" dc:"等级名称"`
	Icon       string `json:"icon" dc:"等级图标"`
	Experience int64  `json:"experience" dc:"等级经验"`
}
