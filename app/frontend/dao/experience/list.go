package dao_experience

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	CreateTime *gtime.Time `json:"createTime" dc:"获得时间"`
	Type       int         `json:"type" dc:"经验类型"`
	Experience int64       `json:"experience" dc:"所得经验"`
}
