package auth

import (
	"context"
	"fmt"
	"server/app/common/consts"
	"server/app/common/library/jwt"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_auth "server/app/frontend/dao/auth"
	dto_auth "server/app/frontend/dto/auth"
	"server/internal/dao"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/frame/g"
)

// ProgramLogin implements service.IAuth.
func (s *sAuth) ProgramLogin(ctx context.Context, req *dto_auth.ProgramLogin) (res *dao_auth.LoginRes, err error) {

	if req.Provider == consts.WEIXIN {
		obj, err := s.weiXin(ctx, req)
		if err != nil {
			return nil, err
		}

		return obj, nil
	}

	return
}

// weiXin implements service.IAuth.
func (s *sAuth) weiXin(ctx context.Context, req *dto_auth.ProgramLogin) (res *dao_auth.LoginRes, err error) {

	options, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatMiniProgramSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	json, err := gjson.DecodeToJson(options)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}

	appid := json.Get("appid")
	secret := json.Get("secret")
	// 请求微信接口
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid, secret, req.Code,
	)
	// 创建一个HTTP客户端请求对象
	resp, err := g.Client().Get(ctx, url)
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION)
	}
	defer resp.Close()

	result, err := gjson.DecodeToJson(resp.ReadAllString())
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION)
	}

	// 判断用户是否注册过
	userDetail, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().WxOpenId, result.Get("openid")).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if userDetail.IsEmpty() {

		entity := g.Map{
			dao.SysUser.Columns().Name:       "新用户" + grand.S(6),
			dao.SysUser.Columns().Birthday:   gtime.NewFromTimeStamp(946656000000 / 1000).UTC(),
			dao.SysUser.Columns().CreateTime: gtime.Now(),
			dao.SysUser.Columns().UpdateTime: gtime.Now(),
			dao.SysUser.Columns().WxOpenId:   result.Get("openid"),
		}

		newSalt := grand.S(6)
		newToken := consts.SYSTEMNAME + grand.S(6) + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity[dao.SysUser.Columns().Salt] = newSalt
		entity[dao.SysUser.Columns().Password] = newToken
		entity[dao.SysUser.Columns().Sex] = 3

		//  获取配置
		userInfo, err := dao.SysConfig.Ctx(ctx).
			Where(dao.SysConfig.Columns().Key, consts.UserSetting).
			Value(dao.SysConfig.Columns().Value)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		userJosn := gjson.New(userInfo)
		//  获取默认头像和封面
		entity[dao.SysUser.Columns().Avatar] = userJosn.Get("avatar")
		entity[dao.SysUser.Columns().Cover] = userJosn.Get("cover")
		entity[dao.SysUser.Columns().Status] = consts.Yes
		entity[dao.SysUser.Columns().LevelId] = userJosn.Get("levelId").Int64()
		exp, err := dao.SysLevel.Ctx(ctx).
			Where(dao.SysLevel.Columns().Id, userJosn.Get("levelId").Int64()).
			Value(dao.SysLevel.Columns().Experience)
		if err != nil {
			return nil, utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity[dao.SysUser.Columns().Experience] = exp.Int()

		rs, err := dao.SysUser.Ctx(ctx).Data(&entity).Insert()
		if err != nil {
			return nil, utils_error.Err(response.DB_SAVE_ERROR)
		}

		rid, err := rs.LastInsertId()
		if err != nil {
			return nil, utils_error.Err(response.DB_SAVE_ERROR)
		}

		token, err := jwt.GenFrontendToken(rid, entity[dao.SysUser.Columns().Name].(string))
		if err != nil {
			return nil, utils_error.Err(response.ACCESS_TOKEN_TIMEOUT)
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

	} else {
		_, err = dao.SysUserLoginLog.Ctx(ctx).Data(g.Map{
			dao.SysUserLoginLog.Columns().Ip:         g.RequestFromCtx(ctx).GetClientIp(),
			dao.SysUserLoginLog.Columns().CreateTime: gtime.Now(),
			dao.SysUserLoginLog.Columns().UserId:     userDetail.GMap().Get(dao.SysUser.Columns().Id).(int64),
		}).Insert()
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}

		token, err := jwt.GenFrontendToken(userDetail.GMap().Get(dao.SysUser.Columns().Id).(int64), userDetail.GMap().Get(dao.SysUser.Columns().Name).(string))
		if err != nil {
			return nil, utils_error.Err(response.ACCESS_TOKEN_TIMEOUT)
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
