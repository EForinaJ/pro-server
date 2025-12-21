package dao_title

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id             int64       `json:"id"     dc:"头衔编号"`
	Name           string      `json:"name"   dc:"头衔名称"`
	Pic            string      `json:"pic"   dc:"头衔封面"`
	Game           string      `json:"game"   dc:"所属分类"`
	ServicePercent int         `json:"servicePercent"  dc:"服务百分比"`
	CreateTime     *gtime.Time `json:"createTime" dc:"创建时间"`
}

type OptionsList struct {
	Id   int64  `json:"id" dc:"用户ID"`
	Name string `json:"name" dc:"昵称"`
}
