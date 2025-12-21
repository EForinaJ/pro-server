package menu

import (
	"context"
	dto_meun "server/app/admin/dto/meun"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCreate implements service.IMenu.
func (s *sMenu) CheckCreate(ctx context.Context, req *dto_meun.Create) (err error) {
	exist, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Path, req.Path).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "菜单路径已存在")
	}

	return
}
