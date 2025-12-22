package dto_distribute

type Create struct {
	Code     string `p:"code" v:"required#请输入订单编号" dc:"订单编号"`
	WitkeyId int64  `p:"witkeyId" v:"required#请设置威客ID" dc:"打手威客ID"`
}
