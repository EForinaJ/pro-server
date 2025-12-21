package level

import (
	"context"
	dto_level "server/app/admin/dto/level"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckEdit implements service.ILevel.
func (s *sLevel) CheckEdit(ctx context.Context, req *dto_level.Edit) (err error) {
	res, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Name, req.Name).
		WhereNotIn(dao.SysLevel.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if res {
		return utils_error.ErrMessage(response.FAILD, "该等级以存在")
	}

	return
}
