package menu

import (
	"context"
	dao_menu "server/app/admin/dao/menu"
	"server/app/common/consts"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IMenu.
func (s *sMenu) GetEdit(ctx context.Context, id int64) (res *dao_menu.Edit, err error) {
	var entity *entity.SysMenu

	err = dao.SysMenu.Ctx(ctx).WherePri(id).Scan(&entity)
	if err != nil {
		return nil, err
	}
	// 布尔值映射优化
	boolMapping := map[int]bool{
		consts.Yes: true,
		consts.Not: false,
	}

	err = gconv.Scan(entity, &res)
	if err != nil {
		return nil, err
	}
	res.IsEnable = boolMapping[entity.IsEnable]
	res.IsMenu = boolMapping[entity.IsMenu]
	res.KeepAlive = boolMapping[entity.KeepAlive]
	res.IsHide = boolMapping[entity.IsHide]
	res.IsHideTab = boolMapping[entity.IsHideTab]
	res.IsIframe = boolMapping[entity.IsIframe]
	res.ShowBadge = boolMapping[entity.ShowBadge]
	res.ShowTextBadge = boolMapping[entity.ShowTextBadge]
	res.FixedTab = boolMapping[entity.FixedTab]
	res.IsFullPage = boolMapping[entity.IsFullPage]

	return
}
