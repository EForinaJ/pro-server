package v1

import (
	dao_category "server/app/frontend/dao/category"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/category/all" method:"get" tags:"分类" summary:"分类列表"`
}
type GetAllRes struct {
	List []*dao_category.List `json:"list" dc:"分类列表"`
}
