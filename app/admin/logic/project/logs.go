package project

import (
	"context"
	dao_project "server/app/admin/dao/project"
	dto_project "server/app/admin/dto/project"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetLogs implements service.IProject.
func (s *sProject) GetLogs(ctx context.Context, req *dto_project.Log) (total int, res []*dao_project.Log, err error) {
	m := dao.SysProjectLog.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysProjectLog.Columns().ProjectId, req.Id).
		OrderDesc(dao.SysProjectLog.Columns().CreateTime)

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysProjectLog
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_project.Log, len(list))
	for i, v := range list {
		var entity *dao_project.Log
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
