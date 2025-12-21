package menu

import (
	"context"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

// AddRelated implements service.IMenu.
func (s *sMenu) AddRelated(ctx context.Context, id int64, menus []int64) (err error) {
	_, err = dao.SysRoleMenu.Ctx(ctx).
		Where(dao.SysRoleMenu.Columns().RoleId, id).Delete()
	if err != nil {
		return err
	}
	entites := make([]*do.SysRoleMenu, len(menus))
	for i, v := range menus {
		entites[i] = &do.SysRoleMenu{
			MenuId: v,
			RoleId: id,
		}
	}

	_, err = dao.SysRoleMenu.Ctx(ctx).
		Data(&entites).
		Insert()
	if err != nil {
		return err
	}

	return
}

// GetRelatedId implements service.IMenu.
func (s *sMenu) GetRelatedId(ctx context.Context, id int64) (ids []int64, err error) {
	Ids, err := dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, id).
		Fields(dao.SysRoleMenu.Columns().MenuId).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	ids = gconv.Int64s(Ids)
	return
}

func init() {
	service.RegisterMenu(&sMenu{})
}
