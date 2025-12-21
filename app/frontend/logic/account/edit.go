package account

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IAccount.
func (s *sAccount) Edit(ctx context.Context, req *dto_account.Edit) (err error) {

	entity := g.Map{
		dao.SysUser.Columns().Name:        req.Name,
		dao.SysUser.Columns().Address:     req.Address,
		dao.SysUser.Columns().Birthday:    gtime.NewFromTimeStamp(req.Birthday / 1000).UTC(),
		dao.SysUser.Columns().Avatar:      req.Avatar,
		dao.SysUser.Columns().Sex:         req.Sex,
		dao.SysUser.Columns().Description: req.Description,
		dao.SysUser.Columns().UpdateTime:  gtime.Now(),
	}

	_, err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = g.Redis().Del(ctx, consts.FrontendAccout+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
