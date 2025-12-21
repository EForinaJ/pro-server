package auth

import (
	"context"
	"server/app/common/consts"
	"server/app/common/library/jwt"
	utils_error "server/app/common/utils/error"
	utlis_lock "server/app/common/utils/lock"
	"server/app/common/utils/response"
	dao_auth "server/app/frontend/dao/auth"
	dto_auth "server/app/frontend/dto/auth"
	"server/internal/dao"
	"server/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Login implements service.IAuth.
func (s *sAuth) Login(ctx context.Context, req *dto_auth.Login) (res *dao_auth.LoginRes, err error) {
	obj := dao.SysUser.Ctx(ctx).Fields(
		dao.SysUser.Columns().Salt,
		dao.SysUser.Columns().Password,
		dao.SysUser.Columns().Id,
	)
	obj = obj.Where(dao.SysUser.Columns().Phone, req.Phone)
	userObj, err := obj.One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	if userObj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}
	var user entity.SysUser
	userObj.Struct(&user)
	randPwd := consts.SYSTEMNAME + req.PassWord + user.Salt
	randPwd = gmd5.MustEncryptString(randPwd)

	if !strings.EqualFold(user.Password, randPwd) {
		// 设置密码错误次数
		errTimes, _ := utlis_lock.SetCount(ctx, consts.LoginCount+req.Phone,
			consts.LoginLock+req.Phone, 1800, 5)
		having := 5 - errTimes
		if having == 0 {
			_, err = g.Redis().Del(ctx, consts.LoginCount+req.Phone)
			if err != nil {
				return nil, utils_error.Err(response.CACHE_READ_ERROR)
			}
			return nil, utils_error.ErrMessage(response.LOGIN_ERROR, "账号已锁定，请30分钟后再试")
		} else {
			return nil, utils_error.ErrMessage(response.LOGIN_ERROR, "密码不正确"+",还有"+gconv.String(having)+"次之后账号将锁定")
		}
	} else {
		_, err = dao.SysUserLoginLog.Ctx(ctx).Data(g.Map{
			dao.SysUserLoginLog.Columns().Ip:         g.RequestFromCtx(ctx).GetClientIp(),
			dao.SysUserLoginLog.Columns().CreateTime: gtime.Now(),
			dao.SysUserLoginLog.Columns().UserId:     user.Id,
		}).Insert()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}

		token, err := jwt.GenFrontendToken(user.Id, user.Name)
		if err != nil {
			return nil, utils_error.Err(response.FAILD)
		}

		expire, err := g.Cfg().Get(ctx, "jwt.frontend.expire")
		if err != nil {
			return nil, utils_error.Err(response.FAILD)
		}
		loginRes := dao_auth.LoginRes{
			Token:  token,
			Expire: gconv.Int(expire),
		}
		res = &loginRes
	}
	return
}
