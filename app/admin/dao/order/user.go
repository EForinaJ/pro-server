package dao_order

type User struct {
	Id     int64  `json:"id"     dc:"编号"`
	Name   string `json:"name" dc:"名称"`
	Avatar string `json:"avatar" dc:"头像"`
}
