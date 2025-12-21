package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id         int64       `json:"id" dc:"ID"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	User       *User       `json:"user" dc:"所属用户"`
	Album      []string    `json:"album" dc:"相册"`
	Commission float64     `json:"commission" dc:"佣金"`
	Rate       int         `json:"rate" dc:"评分"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
