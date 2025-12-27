package withdraw

import (
	"context"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

type sWithdraw struct{}

// CheckStatusNotApply implements service.IWithdraw.
func (s *sWithdraw) CheckStatusNotApply(ctx context.Context, id int64) (res bool, err error) {
	status, err := dao.SysWithdraw.Ctx(ctx).Where(dao.SysWithdraw.Columns().Id, id).Value(dao.SysWithdraw.Columns().Status)
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	if gconv.Int(status) == consts.StatusApply {
		return true, nil
	}

	return
}

func init() {
	service.RegisterWithdraw(&sWithdraw{})
}
