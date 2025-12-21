package user

import (
	"context"
	dao_user "server/app/admin/dao/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IUser.
func (s *sUser) GetDetail(ctx context.Context, id int64) (res *dao_user.Detail, err error) {
	info, err := dao.SysUser.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}
	// 2. 使用结构体转换代替手动映射
	var detail *dao_user.Detail
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	levelModel, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Id, info.GMap().Get(dao.SysUser.Columns().LevelId)).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var levelDetail *dao_user.Level
	if err := gconv.Scan(levelModel.Map(), &levelDetail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Level = levelDetail

	return detail, nil
}
