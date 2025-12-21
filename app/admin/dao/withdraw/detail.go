package dao_withdraw

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id            int64       `json:"id" dc:"ID"`
	Code          string      `json:"code" dc:"UID"`
	Manage        string      `json:"manage" dc:"管理昵称"`
	Witkey        string      `json:"witkey" dc:"威客昵称"`
	Name          string      `json:"name" dc:"真是姓名"`
	Money         float64     `json:"money" dc:"提现金额"`
	SettledAmount float64     `json:"settledAmount" dc:"实际到账金额"`
	ServiceFee    float64     `json:"serviceFee" dc:"平台服务费"`
	Mode          int         `json:"mode" dc:"提现平台：1支付宝，2微信"`
	Number        string      `json:"number" dc:"提现账户"`
	ReceiptFiles  []string    `json:"receiptFiles" dc:"提现平台回执收据文件"`
	ReceiptNum    string      `json:"receiptNum" dc:"提现平台回执收据号"`
	Remark        string      `json:"remark" dc:"备注"`
	Status        int         `json:"status" dc:"状态"`
	CreateTime    *gtime.Time `json:"createTime" dc:"创建时间"`
}
