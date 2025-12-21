package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_account "server/app/frontend/dao/account"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetWithdrawList implements service.IAccount.
func (s *sAccount) GetWithdrawList(ctx context.Context, req *dto_account.WithdrawQuery) (total int, res []*dao_account.WithdrawList, err error) {
	m := dao.SysWithdraw.Ctx(ctx).
		Where(dao.SysWithdraw.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysWithdraw.Columns().CreateTime)

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysWithdraw
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_account.WithdrawList, len(list))
	for i, v := range list {
		var entity *dao_account.WithdrawList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
