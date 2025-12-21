package site

import (
	"context"
	dao_site "server/app/admin/dao/site"
	dao_system "server/app/admin/dao/system"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSite struct{}

// GetInfo implements service.ISite.
func (s *sSite) GetInfo(ctx context.Context) (res *dao_site.Detail, err error) {
	options, err := g.Redis().Get(ctx, "site")
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

	withdrawSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var pay *dao_system.Withdraw
	err = withdrawSetting.Scan(&pay)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	var site *dao_site.Detail
	err = gconv.Scan(base, &site)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	site.Symbol = pay.Symbol
	site.FileSize = file.FileSize
	site.FileType = file.FileType
	site.ImageSize = file.ImageSize
	site.ImageType = file.ImageType
	res = site
	err = g.Redis().SetEX(ctx, "site", site, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR)
	}
	return
}

func init() {
	service.RegisterSite(&sSite{})
}
