package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type FundLog struct {
	Id         int64       `json:"id" dc:"ID"`
	Type       int         `json:"type" dc:"类型"`
	Mode       int         `json:"mode" dc:"组别方式"`
	After      float64     `json:"after" dc:"变动前"`
	Before     float64     `json:"before" dc:"变动后"`
	Money      float64     `json:"money" dc:"变动数量"`
	Remark     string      `json:"remark" dc:"备注"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}

type Log struct {
	Name       string      `json:"name" dc:"操作者"`
	CreateTime *gtime.Time `json:"createTime" dc:"操作时间"`
	Mode       int         `json:"mode" dc:"操作内容"`
	Content    string      `json:"content" dc:"详细内容"`
}
