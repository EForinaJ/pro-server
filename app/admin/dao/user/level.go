package dao_user

type Level struct {
	Id   int64  `json:"id" dc:"ID"`
	Name string `json:"name" dc:"等级"`
	Icon string `json:"icon" dc:"图标"`
}
