package recharge

import (
	"context"
	"time"

	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	dao_recharge "server/app/frontend/dao/recharge"
	dto_recharge "server/app/frontend/dto/recharge"
	"server/internal/dao"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// WechatMiniProgram implements service.IRecharge.
func (s *sRecharge) WechatMiniProgram(ctx context.Context, req *dto_recharge.WechatMiniProgram) (res *dao_recharge.WechatMiniProgram, err error) {

	code := utils_snow.GetCode(ctx, consts.CZ)
	entity := g.Map{
		dao.SysRecharge.Columns().Code:       code,
		dao.SysRecharge.Columns().PayType:    consts.RechargeTypeWechat,
		dao.SysRecharge.Columns().Money:      req.Money,
		dao.SysRecharge.Columns().UserId:     ctx.Value("userId"),
		dao.SysRecharge.Columns().Status:     consts.PayStatusPending,
		dao.SysRecharge.Columns().CreateTime: gtime.Now(),
	}
	rs, err := dao.SysRecharge.Ctx(ctx).
		Data(entity).Insert()
	if err != nil {
		return nil, utils_error.Err(response.DB_SAVE_ERROR)
	}
	rid, err := rs.LastInsertId()
	if err != nil {
		return nil, utils_error.Err(response.DB_SAVE_ERROR)
	}

	wechatPaySetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatPaySetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	wechatPaySettingJson, err := gjson.DecodeToJson(wechatPaySetting)
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID)
	}
	mchId := wechatPaySettingJson.Get("mchId").String()
	mchKey := wechatPaySettingJson.Get("mchKey").String()
	serialNo := wechatPaySettingJson.Get("serialNo").String()
	apiclientKey := "." + wechatPaySettingJson.Get("apiclientKey").String()

	client, err := wechat.NewClientV3(mchId, serialNo, mchKey, gfile.GetContents(apiclientKey))
	if err != nil {
		return nil, utils_error.Err(response.PAY_ERROR)
	}

	openId, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).Value(dao.SysUser.Columns().WxOpenId)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	wechatMiniProgramSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatMiniProgramSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	wechatMiniProgramJson, err := gjson.DecodeToJson(wechatMiniProgramSetting)
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID)
	}

	appId := wechatMiniProgramJson.Get("appId")
	money := gconv.Float64(decimal.NewFromFloat(req.Money).Mul(decimal.NewFromFloat(100)))

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	notify, err := g.Cfg().Get(ctx, "pay.recharge.notify")
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID)
	}
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", appId.String()).
		Set("description", "充值").
		Set("out_trade_no", code).
		Set("time_expire", expire).
		Set("notify_url", notify).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", money).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", openId.String())
		})

	wxRsp, err := client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		return nil, utils_error.Err(response.PAY_ERROR)
	}
	if wxRsp.Code != wechat.Success {
		return nil, utils_error.Err(response.PAY_ERROR)
	}

	// 小程序
	applet, err := client.PaySignOfApplet(appId.String(), wxRsp.Response.PrepayId)
	if err != nil {
		return nil, utils_error.Err(response.PAY_ERROR)
	}

	err = gconv.Scan(applet, &res)
	if err != nil {
		return nil, utils_error.Err(response.PAY_ERROR)
	}

	res.Id = rid

	return
}
