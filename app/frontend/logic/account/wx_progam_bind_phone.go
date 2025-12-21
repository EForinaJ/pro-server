package account

import (
	"context"
	"fmt"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// WxProgramBindPhone implements service.IAccount.
func (s *sAccount) WxProgramBindPhone(ctx context.Context, req *dto_account.WxProgramBindPhone) (res string, err error) {
	phone, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, ctx.Value("userId")).Value(dao.SysUser.Columns().Phone)
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR)
	}
	if phone.String() != "" {
		return "", utils_error.ErrMessage(response.UPDATE_FAILED, "用户已绑定过手机号")
	}

	options, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatMiniProgramSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR)
	}

	json, err := gjson.DecodeToJson(options)
	if err != nil {
		return "", utils_error.Err(response.CACHE_READ_ERROR)
	}

	appid := json.Get("appid")
	secret := json.Get("secret")
	// 请求微信access_token
	tokenUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		appid, secret,
	)
	// 创建一个HTTP客户端请求对象
	resp, err := g.Client().Get(ctx, tokenUrl)
	if err != nil {
		return "", utils_error.Err(response.EXCEPTION)
	}
	defer resp.Close()

	result, err := gjson.DecodeToJson(resp.ReadAllString())
	if err != nil {
		return "", utils_error.Err(response.PARAM_INVALID)
	}

	// 请求微信phone
	phoneUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s",
		result.Get("access_token"),
	)

	phoneCode := fmt.Sprintf(
		`{"code":"%s"}`,
		req.Code,
	)

	//
	phoneResp, err := g.Client().Post(ctx, phoneUrl, phoneCode)

	if err != nil {
		return "", utils_error.Err(response.EXCEPTION)
	}
	defer phoneResp.Close()

	phoneResult, err := gjson.DecodeToJson(phoneResp.ReadAllString())
	if err != nil {
		return "", utils_error.Err(response.PARAM_INVALID)
	}

	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, ctx.Value("userId")).Data(g.Map{
		dao.SysUser.Columns().Phone: phoneResult.GetJson("phone_info").Get("phoneNumber"),
	}).Update()
	if err != nil {
		return "", utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = g.Redis().Del(ctx, consts.FrontendAccout+gconv.String(ctx.Value("userId")))
	if err != nil {
		return "", utils_error.Err(response.CACHE_READ_ERROR)
	}

	return phoneResult.GetJson("phone_info").Get("phoneNumber").String(), nil
}
