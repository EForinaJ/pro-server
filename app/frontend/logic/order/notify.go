package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	"server/internal/dao"

	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Notify implements service.IOrder.
func (s *sOrder) Notify(ctx context.Context) (err error) {
	// 验证标识
	notifyReq, err := wechat.V3ParseNotify(g.RequestFromCtx(ctx).Request)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID)
	}

	if notifyReq.EventType != "TRANSACTION.SUCCESS" {
		return utils_error.Err(response.PARAM_INVALID)
	}

	// 获取微信平台证书
	wechatPaySetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatPaySetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	wechatPaySettingJson, err := gjson.DecodeToJson(wechatPaySetting)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID)
	}
	mchId := wechatPaySettingJson.Get("mchId").String()
	mchKey := wechatPaySettingJson.Get("mchKey").String()
	serialNo := wechatPaySettingJson.Get("serialNo").String()
	apiclientKey := "." + wechatPaySettingJson.Get("apiclientKey").String()

	client, err := wechat.NewClientV3(mchId, serialNo, mchKey, gfile.GetContents(apiclientKey))
	if err != nil {
		return utils_error.Err(response.PAY_ERROR)
	}

	pubKey := "." + wechatPaySettingJson.Get("pubKey").String()
	pubId := wechatPaySettingJson.Get("pubId").String()

	err = client.AutoVerifySignByPublicKey(gfile.GetBytes(pubKey), pubId)
	if err != nil {
		return utils_error.Err(response.PAY_ERROR)
	}

	// 获取微信平台证书
	certMap := client.WxPublicKeyMap()
	// 验证异步通知的签名
	err = notifyReq.VerifySignByPKMap(certMap)
	if err != nil {
		return utils_error.Err(response.PAY_ERROR)
	}

	// 普通支付通知解密
	result, err := notifyReq.DecryptPayCipherText(mchKey)
	if err != nil {
		return utils_error.Err(response.PAY_ERROR)
	}

	if result.OutTradeNo == "" || result.Amount.Total == 0 {
		return utils_error.Err(response.PAY_ERROR)
	}

	order, err := dao.SysOrder.Ctx(ctx).
		// Where(dao.SysOrder.Columns().PayCode, result.OutTradeNo).
		Fields(
			dao.SysOrder.Columns().Id,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().ActualAmount,
			dao.SysOrder.Columns().PayStatus).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	actualAmount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().ActualAmount)))
	payTotal := decimal.NewFromFloat(gconv.Float64(result.Amount.Total)).Div(decimal.NewFromFloat(100))
	g.Dump(actualAmount, payTotal)
	if !actualAmount.Equal(payTotal) {
		return utils_error.Err(response.PAY_ERROR)
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().PayStatus)) == consts.PayStatusSuccess {
		return utils_error.Err(response.PAY_ERROR)
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().Id)).
		Data(g.Map{
			dao.SysOrder.Columns().PayStatus: consts.PayStatusSuccess,
			dao.SysOrder.Columns().PayTime:   gtime.Now(),
			dao.SysOrder.Columns().PayMode:   consts.PayModeBalance,
			dao.SysOrder.Columns().PayAmount: actualAmount,
			dao.SysOrder.Columns().Status:    consts.OrderStatusPendingOrder,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	billEntity := g.Map{
		dao.SysUserBill.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
		dao.SysUserBill.Columns().RelatedId:  order.GMap().Get(dao.SysOrder.Columns().Id),
		dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
		dao.SysUserBill.Columns().Type:       consts.BillTypeOrder,
		dao.SysUserBill.Columns().Amount:     actualAmount,
		dao.SysUserBill.Columns().Mode:       consts.Sub,
		dao.SysUserBill.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(billEntity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
