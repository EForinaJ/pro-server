package consts

const (
	Add = 1
	Sub = 2
)

const (
	PayModeAliPay           = 1 // 支付宝
	PayModeWechatPay        = 2 // 微信
	PayModeBalance          = 3 // 余额
	PayModePersonalTransfer = 4 // 个人转账
)

const (
	FundLogModeAdd = 1
	FundLogModeSub = 2
)

const (
	SettlementModeDeduction = 1 // 扣款
	SettlementModeFull      = 2 // 全款
)

const (
	JobLogModeEdit       = 1 // 修改派单
	JobLogModeCancel     = 2 // 取消派单
	JobLogModePublish    = 3 // 派发威客
	JobLogModeSettlement = 4 // 结算派单
)
