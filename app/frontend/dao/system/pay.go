package dao_system

type Pay struct {
	Symbol          string  `json:"symbol" dc:"支付符号，表示支付类型或货币类型"`
	RechargeType    []int   `json:"rechargeType" dc:"充值模式，表示支持的充值方式列表"`
	PayType         []int   `json:"payType" dc:"支付模式，表示支持的支付方式列表"`
	MinWithdraw     float64 `json:"minWithdraw" dc:"最小提现金额，表示允许的最小提现数额"`
	WithDrawPercent int     `json:"withDrawPercent" dc:"提现百分比，表示提现时的手续费比例"`
	ServicePercent  int     `json:"servicePercent" dc:"服务费百分比，表示支付时收取的服务费比例"`
}
