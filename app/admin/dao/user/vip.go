package dao_user

type Vip struct {
	Id   int64  `json:"id" dc:"会员ID"`
	Name string `json:"name" dc:"会员名称"`
	Icon string `json:"icon" dc:"会员图标"`
}
