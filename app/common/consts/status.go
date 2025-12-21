package consts

const (
	StatusApply   = 1 // 待审核
	StatusSuccess = 2 // 通过
	StatusFail    = 3 // 拒绝
)

const (
	Yes = 2
	Not = 1
)

const (
	Disable = 1
	Enable  = 2
)

const (
	WithdrawStatusApply  = 1 // 待审核
	WithdrawStatuSuccess = 2 // 提现成功
	WithdrawStatuFail    = 3 // 提现失败
)

const (
	PayStatusPending = 1 // 未支付
	PayStatusSuccess = 2 // 支付成功
	PayStatusRefund  = 3 // 已退款
)

const (
	OrderStatusPendingPayment = 1 // 待付款
	OrderStatusPendingOrder   = 2 // 待服务
	OrderStatusInProgress     = 3 // 进行中
	OrderStatusComplete       = 4 // 已完成
	OrderStatusCancel         = 5 // 已取消
	OrderStatusRefund         = 6 // 已退款
)

const (
	OrderWitkeyStatusPending    = 1 // 待服务
	OrderWitkeyStatusInProgress = 2 // 进行中
	OrderWitkeyStatusComplete   = 3 // 已完成
	OrderWitkeyStatusCancel     = 4 // 已取消
)

const (
	ProjectStatusPending           = 1 // 待服务
	ProjectStatusInProgress        = 2 // 进行中
	ProjectStatusPendingSettlement = 3 // 待结算
	ProjectStatusCancel            = 4 // 已取消
	ProjectStatusSettlement        = 5 // 已结算
)
