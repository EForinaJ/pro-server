package permission

import (
	"context"
	dto_permission "server/app/admin/dto/permission"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IPermission.
func (s *sPermission) Create(ctx context.Context, req *dto_permission.Create) (err error) {
	var entity *do.SysPermission
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	_, err = dao.SysPermission.Ctx(ctx).Insert(&entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
