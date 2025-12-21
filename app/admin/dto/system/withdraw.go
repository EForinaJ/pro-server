package dto_system

type Withdraw struct {
	Symbol          string  `v:"required#请输入货币标识" json:"symbol" p:"symbol" dc:"网站货币标识"`
	MinWithdraw     float64 `json:"minWithdraw" p:"minWithdraw" dc:"最小提现金额，表示允许的最小提现数额" v:"required|min:50#最小提现金额不能为空|最小提现金额必须大于或等于50"`
	WithDrawPercent float64 `json:"withDrawPercent" p:"withDrawPercent" dc:"提现百分比，表示提现时的手续费比例" v:"required|between:0,100#提现百分比不能为空|提现百分比必须在0到100之间"`
}
