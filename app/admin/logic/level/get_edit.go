package level

import (
	"context"
	dao_level "server/app/admin/dao/level"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetEdit implements service.ILevel.
func (s *sLevel) GetEdit(ctx context.Context, id int64) (res *dao_level.Edit, err error) {
	err = dao.SysLevel.Ctx(ctx).Where(dao.SysLevel.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
