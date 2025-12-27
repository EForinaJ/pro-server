package dao_user

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id      int64   `json:"id" dc:"用户ID"`
	Avatar  string  `json:"avatar" dc:"头像"`
	Name    string  `json:"name" dc:"昵称"`
	Sex     int     `json:"sex" dc:"性别"`
	Phone   string  `json:"phone" dc:"手机号"`
	Balance float64 `json:"balance" dc:"余额"`
	Status  int     `json:"status" dc:"状态"`
	// LoginIp         string      `json:"loginIp" dc:"登录IP"`
	// LoginTime       *gtime.Time `json:"loginTime" dc:"登录时间"`
	CreateTime *gtime.Time `json:"createTime" dc:"注册时间"`
}

type OptionsList struct {
	Id   int64  `json:"id" dc:"用户ID"`
	Name string `json:"name" dc:"昵称"`
}

type BalanceList struct {
	Id         int64       `json:"id" dc:"ID"`
	Related    string      `json:"related" dc:"关联订单"`
	Type       int         `json:"type" dc:"变更类型"`
	Mode       int         `json:"mode" dc:"组别方式"`
	After      float64     `json:"after" dc:"变动前"`
	Before     float64     `json:"before" dc:"变动后"`
	Amount     float64     `json:"amount" dc:"变动数量"`
	Remark     string      `json:"remark" dc:"原因"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
