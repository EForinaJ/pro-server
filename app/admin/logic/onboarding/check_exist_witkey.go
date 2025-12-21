package onboarding

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckExistWitkey implements service.IOnboarding.
func (s *sOnboarding) CheckExistWitkey(ctx context.Context, id int64) (res bool, err error) {
	userId, err := dao.SysOnboarding.Ctx(ctx).WherePri(id).Value(dao.SysOnboarding.Columns().UserId)
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}

	exist, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().UserId, userId.Int64()).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}

	return !exist, nil
}
