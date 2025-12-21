package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	dao_order "server/app/frontend/dao/order"
	"server/internal/dao"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// WeChatMiniProgramPay implements service.IOrder.
func (s *sOrder) WeChatMiniProgramPay(ctx context.Context, id int64) (res *dao_order.WechatMiniProgram, err error) {

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

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Fields(dao.SysOrder.Columns().ActualAmount, dao.SysOrder.Columns().Code).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	actualAmount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().ActualAmount)))
	money := gconv.Float64(actualAmount.Mul(decimal.NewFromFloat(100)))
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	notify, err := g.Cfg().Get(ctx, "pay.order.notify")
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID)
	}
	code := utils_snow.GetCode(ctx, "MP")

	_, err = dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		// Data(dao.SysOrder.Columns().PayCode, code).
		Update()
	if err != nil {
		return nil, utils_error.Err(response.DB_SAVE_ERROR)
	}

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", appId.String()).
		Set("description", "订单支付").
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

	res.Id = id

	return
}
