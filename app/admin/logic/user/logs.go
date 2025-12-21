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

// GetLogs implements service.IUser.
func (s *sUser) GetLogs(ctx context.Context, req *dto_user.Log) (total int, res []*dao_user.Log, err error) {
	m := dao.SysUserLog.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysUserLog.Columns().UserId, req.Id).
		OrderDesc(dao.SysUserLog.Columns().CreateTime)

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysUserLog
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_user.Log, len(list))
	for i, v := range list {
		var entity *dao_user.Log
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		//  操作人
		user, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id, v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Name = gconv.String(user)

		res[i] = entity
	}
	return
}
