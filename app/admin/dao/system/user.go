package dao_system

type User struct {
	Avatar string `json:"avatar" dc:"默认头像"`
	// Cover             string `json:"cover" dc:"默认背景"`
	LevelId string `json:"levelId" dc:"默认等级"`
	// PaymentExperience string `json:"paymentExperience" dc:"下单可获得经验值"`
	CheckInExperience string `json:"checkInExperience" dc:"签到可获得经验值"`
}
