package v1

import (
	dao_account "server/app/frontend/dao/account"
	dto_account "server/app/frontend/dto/account"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/account" method:"get" tags:"账户" summary:"获取登录账户信息"`
}
type GetDetailRes struct {
	*dao_account.Detail
}

type WxProgramBindPhoneReq struct {
	g.Meta `path:"/account/wx/program/phone" method:"post" tags:"账户" summary:"微信小程序绑定手机"`
	*dto_account.WxProgramBindPhone
}
type WxProgramBindPhoneRes struct {
	Phone string `json:"phone" dc:"手机号"`
}

type EditReq struct {
	g.Meta `path:"/account/edit" method:"post" tags:"账户" summary:"修改用户"`
	*dto_account.Edit
}
type EditRes struct{}

type ChangePasswordReq struct {
	g.Meta `path:"/account/change/password" method:"post" tags:"账户" summary:"修改密码"`
	*dto_account.Password
}
type ChangePasswordRes struct{}

type SignInReq struct {
	g.Meta `path:"/account/sign/in" method:"post" tags:"账户" summary:"用户签到"`
}
type SignInRes struct {
	Experience int `json:"experience" dc:"经验值"`
}

type GetWithdrawListReq struct {
	g.Meta `path:"/account/withdraw/list" method:"get" tags:"账户" summary:"获取登录提现列表"`
	*dto_account.WithdrawQuery
}
type GetWithdrawListRes struct {
	Total int                         `json:"total" dc:"总数"`
	List  []*dao_account.WithdrawList `json:"list" dc:"列表"`
}

type WithdrawReq struct {
	g.Meta `path:"/account/withdraw" method:"post" tags:"账户" summary:"账户提现申请"`
	*dto_account.Withdraw
}
type WithdrawRes struct{}

// 账户卡包
// ====================================
type GetCardListReq struct {
	g.Meta `path:"/account/card/list" method:"get" tags:"账户" summary:"获取账户提现账户列表"`
}
type GetCardListRes struct {
	List []*dao_account.CardList `json:"list" dc:"现账户列表"`
}

type CreateCardReq struct {
	g.Meta `path:"/account/create/card" method:"post" tags:"账户" summary:"创建提现账户"`
	*dto_account.CreateCard
}
type CreateCardRes struct{}

type DeleteCardReq struct {
	g.Meta `path:"/account/delete/card" method:"post" tags:"账户" summary:"删除提现账户"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type DeleteCardRes struct{}
