package consts

const (
	CredentialTypeIdCrad   = 1 //身份证
	CredentialTypePassport = 2 //护照
)

const (
	WitkeyLogTypeCommission    = 3 // 平台变更佣金
	WitkeyLogTypePunish        = 4 // 处罚用户
	WitkeyLogTypePunishRevoke  = 5 // 处罚撤回
	WitkeyLogTypeMarginDeposit = 6 // 平台变更保证金
)

const (
	OrderLogTypeDistribute  = 1 // 添加项目打手
	OrderLogTypeRefund      = 2 // 手动退款
	OrderLogTypePayment     = 3 // 确认支付
	OrderLogTypeAddDiscount = 4 // 添加优惠
	OrderLogTypeAfterasles  = 5 // 退款售后
	OrderLogTypeCancel      = 6 // 关闭订单
)

const (
	PaymentOrder   = 1 // 订单支付
	UserIncomeGift = 2 // 礼物收入
)

const (
	UserLogTypeDeleteWitkey   = 8 // 删除威客信息
	UserLogTypeEditWitkey     = 7 // 修改威客信息
	UserLogTypeAddWitkey      = 6 // 添加威客信息
	UserLogTypeRechargeRevoke = 5 // 撤回预存
	UserLogTypeBalance        = 4 // 修改用户余额
	UserLogTypeRecharge       = 3 // 充值预存
	UserLogTypeEdit           = 2 // 修改
	UserLogTypeCreate         = 1 // 创建
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
	BillTypeRecharge             = 1 // 账户充值
	BillTypeRefund               = 2 // 商品退款
	BillTypeOrder                = 3 // 购买商品
	BillTypeSettlementCommission = 4 // 结算佣金
)

const (
	RechargeTypeAlyPay = 1 // 支付宝
	RechargeTypeWechat = 2 // 微信
	RechargeTypeManual = 3 // 客服预存
)

const (
	ProjectLogTypeSettlement = 1 // 结算
	ProjectLogTypeCancel     = 2 // 取消
	ProjectLogTypeAftersales = 4 // 售后退款
	ProjectLogTypeCreate     = 3 // 创建
)
const (
	ExperienceTypeCheckIn = 1 // 签到经验
	ExperienceTypePayment = 2 // 下单购买
)
