package dao_level

type Edit struct {
	Id         int64  `json:"id"     dc:"编号"`
	Name       string `json:"name" dc:"等级名称"`
	Level      int    `json:"level" dc:"等级数值"`
	Icon       string `json:"icon" dc:"等级图标"`
	Experience int    `json:"experience" dc:"最大经验值"`
	// Benefits   string      `json:"benefits" dc:"等级权益说明"`
}
