package dto_system

type User struct {
	Avatar  string `v:"required#请输入默认头像" json:"avatar" p:"avatar" dc:"默认头像"`
	LevelId string `v:"required#请选择默认等级" json:"levelId"  p:"levelId" dc:"默认等级"`
	// PaymentExperience string `v:"required#请输入下的那经验值" json:"paymentExperience"  p:"paymentExperience" dc:"下单可获得经验值"`
	CheckInExperience string `v:"required#请输入签到经验值" json:"checkInExperience"  p:"checkInExperience" dc:"签到可获得经验值"`
}
