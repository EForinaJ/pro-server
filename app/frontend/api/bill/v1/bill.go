package v1

import (
	dao_bill "server/app/frontend/dao/bill"
	dto_bill "server/app/frontend/dto/bill"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/bill/list" method:"get" tags:"账单" summary:"获取账单列表"`
	*dto_bill.Query
}
type GetListRes struct {
	Total int              `json:"total" dc:"总数"`
	List  []*dao_bill.List `json:"list" dc:"列表"`
}
