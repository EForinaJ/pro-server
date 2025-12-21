package dao_freelancer

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Uid        int64       `json:"uid" dc:"用户id"`
	Examiner   string      `json:"examiner" dc:"考官"`
	Name       string      `json:"name" dc:"昵称"`
	Title      string      `json:"title" dc:"头衔"`
	Category   string      `json:"category" dc:"所属分类"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
