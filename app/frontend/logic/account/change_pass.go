package account

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// ChangePass implements service.IAccount.
func (s *sAccount) ChangePass(ctx context.Context, req *dto_account.Password) (err error) {

	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.NewPass + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Data(g.Map{
			dao.SysUser.Columns().Salt:       newSalt,
			dao.SysUser.Columns().Password:   newToken,
			dao.SysUser.Columns().UpdateTime: gtime.Now(),
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
