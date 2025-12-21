package auth

import (
	"context"

	dto_auth "server/app/admin/dto/auth"
	service "server/app/admin/service"
	"server/app/common/consts"
	"server/app/common/library/jwt"
	utils_error "server/app/common/utils/error"
	utlis_lock "server/app/common/utils/lock"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAtuh struct{}

// Login implements admin_service.IAuth.
func (s *sAtuh) Login(ctx context.Context, req *dto_auth.Login) (res interface{}, err error) {
	obj := dao.SysManage.Ctx(ctx).Fields(
		dao.SysManage.Columns().Salt,
		dao.SysManage.Columns().Password,
		dao.SysManage.Columns().Id,
	)
	obj = obj.Where(dao.SysManage.Columns().Phone, req.Phone)

	userObj, err := obj.One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	if userObj == nil {
		return nil, gerror.NewCode(gcode.New(response.NOT_FOUND, "", nil), response.CodeMsg(response.NOT_FOUND))
	}
	var user entity.SysManage
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
				return nil, gerror.NewCode(gcode.New(response.CACHE_READ_ERROR, "", nil), response.CodeMsg(response.CACHE_READ_ERROR))
			}
			return nil, gerror.NewCode(gcode.New(response.LOGIN_ERROR, "", nil), "账号已锁定，请30分钟后再试")
		} else {
			return nil, gerror.NewCode(gcode.New(response.LOGIN_ERROR, "", nil), "密码不正确"+",还有"+gconv.String(having)+"次之后账号将锁定")
		}
	} else {
		_, err = dao.SysManageLoginLog.Ctx(ctx).Data(g.Map{
			dao.SysManageLoginLog.Columns().Ip:         g.RequestFromCtx(ctx).GetClientIp(),
			dao.SysManageLoginLog.Columns().CreateTime: gtime.Now(),
			dao.SysManageLoginLog.Columns().ManageId:   user.Id,
		}).Insert()
		if err != nil {
			return "", utils_error.Err(response.DB_READ_ERROR)
		}
		token, err := jwt.GenAdminToken(user.Id, user.Name)
		if err != nil {
			return nil, gerror.NewCode(gcode.New(response.FAILD, "", nil), response.CodeMsg(response.FAILD))
		}
		res = token
	}
	return
}

func init() {
	service.RegisterAuth(&sAtuh{})
}
