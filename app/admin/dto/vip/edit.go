package dto_vip

type Edit struct {
	Id         int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"会员id"`
	Name       string  `p:"name" v:"required#请输入会员名称" dc:"会员名称"`
	Level      int     `p:"level" v:"required#请输入会员等级" dc:"会员等级"`
	Icon       string  `p:"icon" v:"required#请上传会员图标" dc:"会员图标"`
	Price      float64 `p:"price" v:"required#请输入会员价格" dc:"会员价格"`
	Duration   int     `p:"duration" v:"required#请输入有效期(天)" dc:"有效期(天)"`
	Benefits   string  `p:"benefits" v:"required#请输入会员权益说明" dc:"会员权益说明"`
	Discount   float64 `p:"discount" v:"required#请输入折扣率(0-1)" dc:"折扣率(0-1)"`
	PointsRate float64 `p:"pointsRate" v:"required#请输入积分倍率" dc:"积分倍率"`
}
