package dao_account

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id          int16       `json:"id" dc:"uid"`
	Name        string      `json:"name" dc:"昵称"`
	Phone       string      `json:"phone" dc:"手机号"`
	Description string      `json:"description" dc:"描述"`
	Sex         int         `json:"sex" dc:"性别"`
	Avatar      string      `json:"avatar" dc:"头像"`
	Balance     float64     `json:"balance" dc:"余额"`
	CreateTime  *gtime.Time `json:"createTime" dc:"注册时间"`
	Birthday    int64       `json:"birthday" dc:"出生日期"`
	Address     string      `json:"address" dc:"所在地"`
	Level       *Level      `json:"level" dc:"用户等级"`
	IsSign      bool        `json:"isSign" dc:"是否签到了"`
	Experience  int64       `json:"experience" dc:"所得经验"`
	Roles       []string    `json:"roles" dc:"角色"`
}
