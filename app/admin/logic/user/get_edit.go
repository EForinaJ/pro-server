package user

import (
	"context"
	dao_user "server/app/admin/dao/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IUser.
func (s *sUser) GetEdit(ctx context.Context, id int64) (res *dao_user.Edit, err error) {
	info, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var detail *dao_user.Edit
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}

	detail.Birthday = gtime.New(info.Map()[dao.SysUser.Columns().Birthday]).TimestampMilli()
	return detail, nil
}
