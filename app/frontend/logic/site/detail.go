package site

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_site "server/app/frontend/dao/site"
	dao_system "server/app/frontend/dao/system"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetInfo implements service.ISite.
func (s *sSite) GetInfo(ctx context.Context) (res *dao_site.Detail, err error) {
	options, err := g.Redis().Get(ctx, "frontend_site")
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}

	if !options.IsEmpty() {
		json, jerr := gjson.DecodeToJson(options)
		if jerr != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR)
		}
		jerr = json.Scan(&res)
		if jerr != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR)
		}
		return
	}
	baseSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.BaseSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var base *dao_system.Base
	err = baseSetting.Scan(&base)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	fileSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var file *dao_system.File
	err = fileSetting.Scan(&file)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	contactSetting, err := dao.SysConfig.
		Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.ContactSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var contact *dao_site.Contact
	err = contactSetting.Scan(&contact)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	withdrawSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var pay *dao_system.Pay
	err = withdrawSetting.Scan(&pay)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	var site *dao_site.Detail
	err = gconv.Scan(base, &site)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	site.MinFileSize = file.MinFileSize
	site.BigFileSize = file.BigFileSize
	site.MediaType = file.MediaType
	site.Symbol = pay.Symbol
	site.MinWithdraw = pay.MinWithdraw
	site.WithDrawPercent = pay.WithDrawPercent
	site.Contact = contact
	res = site
	err = g.Redis().SetEX(ctx, "frontend_site", site, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR)
	}
	return
}
