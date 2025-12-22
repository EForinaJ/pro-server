package bill

import (
	"context"
	dao_bill "server/app/admin/dao/bill"
	dto_bill "server/app/admin/dto/bill"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IBill.
func (s *sBill) GetList(ctx context.Context, req *dto_bill.Query) (total int, res []*dao_bill.List, err error) {
	m := dao.SysUserBill.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysUserBill.Columns().CreateTime)

	if req.Code != "" {
		m = m.Where(dao.SysUserBill.Columns().Code, req.Code)
	}
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").
			Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysUserBill.Columns().UserId, userId)
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

		//  下单用户
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Fields(dao.SysUser.Columns().Name,
				dao.SysUser.Columns().Avatar).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		entity.User = &dao_bill.User{
			Id:     v.UserId,
			Name:   gconv.String(user.GMap().Get(dao.SysUser.Columns().Name)),
			Avatar: gconv.String(user.GMap().Get(dao.SysUser.Columns().Avatar)),
		}

		entity.RelatedCode, err = getTransactionCode(ctx, v)
		if err != nil {
			return 0, nil, err
		}
		res[i] = entity
	}

	return
}

func getTransactionCode(ctx context.Context, entity *entity.SysUserBill) (code string, err error) {
	switch entity.Type {
	case consts.BillTypeRecharge:
		obj, err := dao.SysRecharge.Ctx(ctx).
			Where(dao.SysRecharge.Columns().Id, entity.RelatedId).
			Value(dao.SysRecharge.Columns().Code)
		if err != nil {
			return "", utils_error.Err(response.DB_READ_ERROR)
		}
		return obj.String(), nil
	case consts.BillTypeRefund:
		obj, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Id, entity.RelatedId).
			Value(dao.SysOrder.Columns().Code)
		if err != nil {
			return "", utils_error.Err(response.DB_READ_ERROR)
		}
		return obj.String(), nil
		// case consts.BillTypeSettlementCommission:
		// 	obj, err := dao.SysProject.Ctx(ctx).
		// 		Where(dao.SysProject.Columns().Id, entity.RelatedId).
		// 		Value(dao.SysProject.Columns().Code)
		// 	if err != nil {
		// 		return "", utils_error.Err(response.DB_READ_ERROR)
		// 	}
		// 	return obj.String(), nil
	}

	return
}
