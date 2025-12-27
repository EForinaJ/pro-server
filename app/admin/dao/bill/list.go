package dao_bill

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id"     dc:"编号"`
	RelatedCode string      `json:"relatedCode"   dc:"关联单号"`
	Code        string      `json:"code"   dc:"账单号"`
	User        *User       `json:"user" dc:"下单用户"`
	Type        int         `json:"type"   dc:"账单类型"`
	Amount      float64     `json:"amount"   dc:"账单金额"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}
