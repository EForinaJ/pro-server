package site

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_site "server/app/frontend/dao/site"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
)

// GetContact implements service.ISite.
func (s *sSite) GetContact(ctx context.Context) (res *dao_site.Contact, err error) {
	contactSetting, err := dao.SysConfig.
		Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.ContactSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	contactJson, err := gjson.DecodeToJson(contactSetting)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var detail *dao_site.Contact
	err = contactJson.Scan(&detail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = detail
	// detail.Platform = contactJson.Get("platform").String()
	return
}
