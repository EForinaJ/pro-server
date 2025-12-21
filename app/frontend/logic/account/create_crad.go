package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateCard implements service.IAccount.
func (s *sAccount) CreateCard(ctx context.Context, req *dto_account.CreateCard) (err error) {
	_, err = dao.SysUserCard.Ctx(ctx).Data(g.Map{
		dao.SysUserCard.Columns().UserId:     ctx.Value("userId"),
		dao.SysUserCard.Columns().Name:       req.Name,
		dao.SysUserCard.Columns().Number:     req.Number,
		dao.SysUserCard.Columns().Type:       req.Type,
		dao.SysUserCard.Columns().NickName:   req.NickName,
		dao.SysUserCard.Columns().CreateTime: gtime.Now(),
		dao.SysUserCard.Columns().UpdateTime: gtime.Now(),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
