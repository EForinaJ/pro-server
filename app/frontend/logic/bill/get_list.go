package bill

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_bill "server/app/frontend/dao/bill"
	dto_bill "server/app/frontend/dto/bill"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IBill.
func (s *sBill) GetList(ctx context.Context, req *dto_bill.Query) (total int, res []*dao_bill.List, err error) {
	m := dao.SysUserBill.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysUserBill.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysUserBill.Columns().CreateTime)

	if req.Type != 0 {
		m = m.Where(dao.SysUserBill.Columns().Type, req.Type)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	var list []*entity.SysUserBill
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = make([]*dao_bill.List, len(list))
	for i, v := range list {
		var entity *dao_bill.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
