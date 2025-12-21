package user

import (
	"context"
	dao_user "server/app/admin/dao/user"
	dto_user "server/app/admin/dto/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetBalanceList implements service.IUser.
func (s *sUser) GetBalanceList(ctx context.Context, req *dto_user.BalanceQuery) (total int, res []*dao_user.BalanceList, err error) {
	m := dao.SysBalance.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysBalance.Columns().CreateTime)
	m = m.WhereIn(dao.SysBalance.Columns().UserId, req.Id)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysBalance
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_user.BalanceList, len(list))
	for i, v := range list {
		var entity *dao_user.BalanceList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		res[i] = entity
	}

	return
}
