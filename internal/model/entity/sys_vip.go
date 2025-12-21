// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysVip is the golang structure for table sys_vip.
type SysVip struct {
	Id         int64       `json:"id"         orm:"id"          description:"会员ID"`     // 会员ID
	Name       string      `json:"name"       orm:"name"        description:"会员名称"`     // 会员名称
	Level      int         `json:"level"      orm:"level"       description:"会员等级"`     // 会员等级
	Icon       string      `json:"icon"       orm:"icon"        description:"会员图标"`     // 会员图标
	Price      float64     `json:"price"      orm:"price"       description:"会员价格"`     // 会员价格
	Duration   int         `json:"duration"   orm:"duration"    description:"有效期(天)"`   // 有效期(天)
	Benefits   string      `json:"benefits"   orm:"benefits"    description:"会员权益说明"`   // 会员权益说明
	Discount   float64     `json:"discount"   orm:"discount"    description:"折扣率(0-1)"` // 折扣率(0-1)
	PointsRate float64     `json:"pointsRate" orm:"points_rate" description:"积分倍率"`     // 积分倍率
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`     // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`     // 更新时间
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`       // 备注
}
