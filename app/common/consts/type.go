package consts

const (
	PaymentOrder   = 1 // 订单支付
	UserIncomeGift = 2 // 礼物收入
)

const (
	UserChangeBalanceTypeSystem   = 1 // 系统变更余额
	UserChangeBalanceTypeRecharge = 2 // 用户充值
)

const (
	WitkeyChangeCommissionTypeSystem     = 1 // 系统变更佣金
	WitkeyChangeCommissionTypeSettlement = 2 // 结算变更佣金
	WitkeyChangeCommissionTypeWithdraw   = 3 // 威客提现变更佣金
)

const (
	BillTypeRecharge             = 1 // 账户充值
	BillTypeRefund               = 2 // 商品退款
	BillTypeOrder                = 3 // 购买商品
	BillTypeSettlementCommission = 3 // 佣金结算
	BillTypeWithdrawCommission   = 4 // 佣金提现
)

const (
	RechargeTypeAlyPay = 1 // 支付宝
	RechargeTypeWechat = 2 // 微信
	RechargeTypeManual = 3 // 客服预存
)

const (
	ExperienceTypeCheckIn = 1 // 签到经验
	ExperienceTypePayment = 2 // 下单购买
)
