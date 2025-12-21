package dao_game

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id"     dc:"分类编号"`
	Name        string      `json:"name"   dc:"分类名称"`
	ShortName   string      `json:"shortName" dc:"短名"`
	Description string      `json:"description"   dc:"分类描述"`
	Pic         string      `json:"pic"   dc:"分类图片"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}
