package account

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckPass implements service.IAccount.
func (s *sAccount) CheckPass(ctx context.Context, req *dto_account.Password) (err error) {
	detail, err := dao.SysUser.Ctx(ctx).Fields(
		dao.SysUser.Columns().Salt,
		dao.SysUser.Columns().Password,
	).Where(dao.SysUser.Columns().Id, ctx.Value("userId")).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if detail.IsEmpty() {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	randPwd := consts.SYSTEMNAME + req.OldPass + gconv.String(detail.GMap().Get(dao.SysUser.Columns().Salt))
	if !strings.EqualFold(gconv.String(detail.GMap().Get(dao.SysUser.Columns().Password)), gmd5.MustEncryptString(randPwd)) {
		return utils_error.ErrMessage(response.FAILD, "原密码错误")
	}
	return
}
