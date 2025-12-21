package v1

import (
	dao_bill "server/app/admin/dao/bill"
	dto_bill "server/app/admin/dto/bill"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/bill/list" method:"get" tags:"账单" summary:"账单列表"`
	*dto_bill.Query
}
type GetListRes struct {
	Total int              `json:"total" dc:"总数"`
	List  []*dao_bill.List `json:"list" dc:"账单列表"`
}
