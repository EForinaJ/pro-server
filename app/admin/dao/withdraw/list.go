package dao_withdraw

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id            int64       `json:"id" dc:"ID"`
	Code          string      `json:"code" dc:"提现编号"`
	Witkey        string      `json:"witkey" dc:"申请威客"`
	Amount        float64     `json:"amount" dc:"提现金额"`
	SettledAmount float64     `json:"settledAmount" dc:"到账金额"`
	ServiceFee    float64     `json:"serviceFee" dc:"平台手续费"`
	Status        int         `json:"status" dc:"状态"`
	Type          int         `json:"type" dc:"到账类型"`
	CreateTime    *gtime.Time `json:"createTime" dc:"创建时间"`
}
