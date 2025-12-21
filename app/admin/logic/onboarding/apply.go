package onboarding

import (
	"context"
	dto_onboarding "server/app/admin/dto/onboarding"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Apply implements service.IOnboarding.
func (s *sOnboarding) Apply(ctx context.Context, req *dto_onboarding.Apply) (err error) {
	switch req.Status {
	case consts.StatusSuccess:
		err = s.success(ctx, req)
		return
	case consts.StatusFail:
		err = s.fail(ctx, req)
		return
	default:
		return utils_error.Err(response.UPDATE_FAILED)
	}
}

func (s *sOnboarding) success(ctx context.Context, req *dto_onboarding.Apply) (err error) {
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
	detail, err := tx.Model(dao.SysOnboarding.Table()).WherePri(req.Id).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	_, err = tx.Model(dao.SysWitkey.Table()).Data(g.Map{
		dao.SysWitkey.Columns().UserId:     detail.GMap().Get(dao.SysOnboarding.Columns().UserId),
		dao.SysWitkey.Columns().GameId:     detail.GMap().Get(dao.SysOnboarding.Columns().GameId),
		dao.SysWitkey.Columns().TitleId:    detail.GMap().Get(dao.SysOnboarding.Columns().TitleId),
		dao.SysWitkey.Columns().CreateTime: gtime.Now(),
		dao.SysWitkey.Columns().UpdateTime: gtime.Now(),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	_, err = tx.Model(dao.SysOnboarding.Table()).Data(g.Map{
		dao.SysOnboarding.Columns().ManageId: ctx.Value("userId"),
		dao.SysOnboarding.Columns().Status:   req.Status,
	}).WherePri(req.Id).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

func (s *sOnboarding) fail(ctx context.Context, req *dto_onboarding.Apply) (err error) {
	_, err = dao.SysOnboarding.Ctx(ctx).Data(g.Map{
		dao.SysOnboarding.Columns().ManageId:     ctx.Value("userId"),
		dao.SysOnboarding.Columns().Status:       req.Status,
		dao.SysOnboarding.Columns().RefuseReason: req.RefuseReason,
	}).WherePri(req.Id).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
