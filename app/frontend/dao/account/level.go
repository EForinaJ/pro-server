package dao_account

type Level struct {
	Id   int16  `json:"id" dc:"uid"`
	Name string `json:"name" dc:"等级名称"`
	Icon string `json:"icon" dc:"等级图标"`
}
