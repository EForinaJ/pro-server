package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	User       *User       `json:"user" dc:"所属用户"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Commission float64     `json:"commission" dc:"佣金"`
	Rate       int         `json:"rate" dc:"评分"`
	CreateTime *gtime.Time `json:"createTime" dc:"入驻时间"`
}

type User struct {
	Name   string `json:"name" dc:"用户名称"`
	Avatar string `json:"avatar" dc:"用户头像"`
	Phone  string `json:"phone" dc:"用户手机号"`
}

type CommissionList struct {
	Id         int64       `json:"id" dc:"ID"`
	Related    string      `json:"related" dc:"关联订单"`
	Type       int         `json:"type" dc:"变更类型"`
	Mode       int         `json:"mode" dc:"组别方式"`
	After      float64     `json:"after" dc:"变动前"`
	Before     float64     `json:"before" dc:"变动后"`
	Money      float64     `json:"money" dc:"变动数量"`
	Remark     string      `json:"remark" dc:"原因"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}

type PunishList struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"处理编号"`
	Manage     string      `json:"manage" dc:"操作人"`
	Money      float64     `json:"money" dc:"处理金额"`   //
	Content    string      `json:"content" dc:"处理内容"` //
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
