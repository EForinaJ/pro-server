package recharge

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

// Notify implements service.IRecharge.
func (s *sRecharge) Notify(ctx context.Context) (err error) {
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

	recharge, err := dao.SysRecharge.Ctx(ctx).
		Where(dao.SysRecharge.Columns().Code, result.OutTradeNo).
		Fields(
			dao.SysRecharge.Columns().Id,
			dao.SysRecharge.Columns().UserId,
			dao.SysRecharge.Columns().Amount,
			dao.SysRecharge.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	orderMoney := decimal.NewFromFloat(gconv.Float64(recharge.GMap().Get(dao.SysRecharge.Columns().Amount)))
	payTotal := decimal.NewFromFloat(gconv.Float64(result.Amount.Total)).Div(decimal.NewFromFloat(100))
	if !orderMoney.Equal(payTotal) {
		return utils_error.Err(response.PAY_ERROR)
	}
	if gconv.Int(recharge.GMap().Get(dao.SysRecharge.Columns().Status)) == consts.PayStatusSuccess {
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

	userBalance, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, recharge.GMap().Get(dao.SysRecharge.Columns().UserId)).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	balance := decimal.NewFromFloat(userBalance.Float64())
	newBalance := balance.Add(orderMoney)
	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, recharge.GMap().Get(dao.SysRecharge.Columns().UserId)).
		Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysRecharge.Table()).
		Where(dao.SysRecharge.Columns().Id, recharge.GMap().Get(dao.SysRecharge.Columns().Id)).
		Data(g.Map{
			dao.SysRecharge.Columns().Status: consts.PayStatusSuccess,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	billEntity := g.Map{
		dao.SysUserBill.Columns().UserId:     recharge.GMap().Get(dao.SysRecharge.Columns().UserId),
		dao.SysUserBill.Columns().RelatedId:  recharge.GMap().Get(dao.SysRecharge.Columns().Id),
		dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
		dao.SysUserBill.Columns().Type:       consts.BillTypeRecharge,
		dao.SysUserBill.Columns().Amount:     orderMoney,
		dao.SysUserBill.Columns().Mode:       consts.Add,
		dao.SysUserBill.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(billEntity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = g.Redis().Del(ctx, consts.FrontendAccout+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
