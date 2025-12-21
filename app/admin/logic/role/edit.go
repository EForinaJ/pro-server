package role

import (
	"context"
	dto_role "server/app/admin/dto/role"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IRole.
func (s *sRole) Edit(ctx context.Context, req *dto_role.Edit) (err error) {
	//  更新数据
	var entity *do.SysRole
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysRole.Ctx(ctx).WherePri(req.Id).
		Data(&entity).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
