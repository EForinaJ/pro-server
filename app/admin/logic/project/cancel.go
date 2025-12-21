package project

import (
	"context"
	dto_project "server/app/admin/dto/project"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cancel implements service.IProject.
func (s *sProject) Cancel(ctx context.Context, req *dto_project.Cancel) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysProject.Table()).
		Where(dao.SysProject.Columns().Id, req.Id).
		Update(g.Map{
			dao.SysProject.Columns().Status: consts.ProjectStatusCancel,
		})
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	// 添加项目日志
	_, err = tx.Model(dao.SysProjectLog.Table()).Data(g.Map{
		dao.SysProjectLog.Columns().ProjectId:  req.Id,
		dao.SysProjectLog.Columns().CreateTime: gtime.Now(),
		dao.SysProjectLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysProjectLog.Columns().Type:       consts.ProjectLogTypeCancel,
		dao.SysProjectLog.Columns().Content:    req.Content,
	}).Insert()
	if err != nil {
		g.Dump(err.Error())
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
