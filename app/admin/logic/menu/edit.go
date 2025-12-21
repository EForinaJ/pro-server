package menu

import (
	"context"
	dto_meun "server/app/admin/dto/meun"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IMenu.
func (s *sMenu) Edit(ctx context.Context, req *dto_meun.Edit) (err error) {
	var entity *do.SysMenu
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	// 布尔值映射优化
	boolMapping := map[bool]int{
		true:  consts.Yes,
		false: consts.Not,
	}
	entity.IsEnable = boolMapping[req.IsEnable]
	entity.IsMenu = boolMapping[req.IsMenu]
	entity.KeepAlive = boolMapping[req.KeepAlive]
	entity.IsHide = boolMapping[req.IsHide]
	entity.IsHideTab = boolMapping[req.IsHideTab]
	entity.IsIframe = boolMapping[req.IsIframe]
	entity.ShowBadge = boolMapping[req.ShowBadge]
	entity.ShowTextBadge = boolMapping[req.ShowTextBadge]
	entity.FixedTab = boolMapping[req.FixedTab]
	entity.IsFullPage = boolMapping[req.IsFullPage]
	_, err = dao.SysMenu.Ctx(ctx).
		WherePri(req.Id).
		Data(&entity).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
