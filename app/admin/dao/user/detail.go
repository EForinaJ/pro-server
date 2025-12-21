package dao_user

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id          int64       `json:"id" dc:"用户ID"`
	Name        string      `json:"name" dc:"用户昵称"`
	Phone       string      `json:"phone" dc:"用户手机号"`
	Sex         int         `json:"sex" dc:"用户性别（1男 2女 3未知）"`
	Address     []string    `json:"address" dc:"用户地址"`
	Birthday    *gtime.Time `json:"birthday" dc:"出生日期"`
	Description string      `json:"description" dc:"用户描述"`
	Avatar      string      `json:"avatar" dc:"头像地址"`
	Balance     float64     `json:"balance" dc:"余额"`
	Deposit     float64     `json:"deposit" dc:"押金"`
	Level       *Level      `json:"level" dc:"等级信息"`
	Status      int         `json:"status" dc:"帐号状态（1停用,2正常）"`
	CreateTime  *gtime.Time `json:"createTime" dc:"注册时间"`
}
