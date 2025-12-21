package dao_account

import "github.com/gogf/gf/v2/os/gtime"

type WithdrawList struct {
	Id         int64       `json:"id" dc:"编号Id"`
	Code       string      `json:"code" dc:"提现码"`
	Money      float64     `json:"money" dc:"账户号码"`
	Type       int         `json:"type" dc:"账户类型"`
	CreateTime *gtime.Time `json:"createTime" dc:"申请时间"`
}

type CardList struct {
	Id       int64  `json:"id" dc:"卡包ID"`
	Name     string `json:"name" dc:"账户名称"`
	Number   string `json:"number" dc:"账户号码"`
	Type     int    `json:"type" dc:"账户类型"`
	NickName string `json:"nickName" dc:"账户昵称"`
}
