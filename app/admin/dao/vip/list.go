package dao_vip

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"会员ID"`
	Name       string      `json:"name" dc:"会员名称"`
	Level      int         `json:"level" dc:"会员等级"`
	Icon       string      `json:"icon" dc:"会员图标"`
	Price      float64     `json:"price" dc:"会员价格"`
	Duration   int         `json:"duration" dc:"有效期(天)"`
	Benefits   string      `json:"benefits" dc:"会员权益说明"`
	Discount   float64     `json:"discount" dc:"折扣率(0-1)"`
	PointsRate float64     `json:"pointsRate" dc:"积分倍率"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
	Remark     string      `json:"remark" dc:"备注"`
}
