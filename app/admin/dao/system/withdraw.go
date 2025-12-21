package dao_system

type Withdraw struct {
	Symbol          string  `json:"symbol" dc:"支付符号，表示支付类型或货币类型"`
	MinWithdraw     float64 `json:"minWithdraw" dc:"最小提现金额，表示允许的最小提现数额"`
	WithDrawPercent int     `json:"withDrawPercent" dc:"提现百分比，表示提现时的手续费比例"`
}
