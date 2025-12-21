// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysVip is the golang structure of table sys_vip for DAO operations like Where/Data.
type SysVip struct {
	g.Meta     `orm:"table:sys_vip, do:true"`
	Id         interface{} // 会员ID
	Name       interface{} // 会员名称
	Level      interface{} // 会员等级
	Icon       interface{} // 会员图标
	Price      interface{} // 会员价格
	Duration   interface{} // 有效期(天)
	Benefits   interface{} // 会员权益说明
	Discount   interface{} // 折扣率(0-1)
	PointsRate interface{} // 积分倍率
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Remark     interface{} // 备注
}
