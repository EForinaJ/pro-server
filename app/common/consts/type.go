package consts

const (
	PaymentOrder   = 1 // 订单支付
	UserIncomeGift = 2 // 礼物收入
)

const (
	UserChangeBalanceTypeSystem   = 1 // 系统变更余额
	UserChangeBalanceTypeRecharge = 2 // 用户充值
	// UserFundLogTypeOrderRefund         = 3  // 订单退款
	// UserFundLogTypeOrderPayMent        = 4  // 订单收款
	// UserFundLogTypeRevokeRecharge      = 5  // 撤回预存
	// UserFundLogTypeWithdraw            = 6  // 提现扣除佣金
	// UserFundLogTypeSettlement          = 7  // 项目结算添加佣金
	// UserFundLogTypeChangeCommission    = 8  // 平台变更佣金
	// UserFundLogTypePunish              = 9  // 处罚扣款佣金
	// UserFundLogTypeRevokePunish        = 10 // 撤回处罚佣金
	// UserFundLogTypeChangeMarginDeposit = 11 // 平台变更保证金
)

const (
	UserChangeCommissionTypeSystem = 1 // 系统变更佣金
	// UserChangeBalanceTypeRecharge = 2 // 用户充值
	// UserFundLogTypeOrderRefund         = 3  // 订单退款
	// UserFundLogTypeOrderPayMent        = 4  // 订单收款
	// UserFundLogTypeRevokeRecharge      = 5  // 撤回预存
	// UserFundLogTypeWithdraw            = 6  // 提现扣除佣金
	// UserFundLogTypeSettlement          = 7  // 项目结算添加佣金
	// UserFundLogTypeChangeCommission    = 8  // 平台变更佣金
	// UserFundLogTypePunish              = 9  // 处罚扣款佣金
	// UserFundLogTypeRevokePunish        = 10 // 撤回处罚佣金
	// UserFundLogTypeChangeMarginDeposit = 11 // 平台变更保证金
)

const (
	BillTypeRecharge = 1 // 账户充值
	BillTypeRefund   = 2 // 商品退款
	BillTypeOrder    = 3 // 购买商品
	// BillTypeSettlementCommission = 4 // 结算佣金
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
