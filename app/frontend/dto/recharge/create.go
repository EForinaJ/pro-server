package dto_recharge

type WechatMiniProgram struct {
	Money float64 `p:"money" v:"required|min:1#请输入充值金额|金额最小为1" dc:"充值金额"`
}
