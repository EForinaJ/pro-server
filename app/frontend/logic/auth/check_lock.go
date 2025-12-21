package auth

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	utlis_lock "server/app/common/utils/lock"
	"server/app/common/utils/response"
	dto_auth "server/app/frontend/dto/auth"
)

// CheckLock implements service.IAuth.
func (s *sAuth) CheckLock(ctx context.Context, req *dto_auth.Login) (err error) {
	if utlis_lock.CheckLock(ctx, consts.LoginLock+req.Phone) {
		return utils_error.ErrMessage(response.FAILD, "账号已锁定，请30分钟后再试")
	}
	return
}
