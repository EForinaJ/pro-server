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

// GetRechargeList implements service.IUser.
func (s *sUser) GetRechargeList(ctx context.Context, req *dto_user.RechargeQuery) (total int, res []*dao_user.RechargeList, err error) {
	m := dao.SysRecharge.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysRecharge.Columns().UserId, req.Id).
		OrderDesc(dao.SysRecharge.Columns().CreateTime)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysRecharge
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_user.RechargeList, len(list))
	for i, v := range list {
		var entity *dao_user.RechargeList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		res[i] = entity
	}
	return
}
