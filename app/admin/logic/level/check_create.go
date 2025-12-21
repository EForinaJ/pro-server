package level

import (
	"context"
	dto_level "server/app/admin/dto/level"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCreate implements service.ILevel.
func (s *sLevel) CheckCreate(ctx context.Context, req *dto_level.Create) (err error) {
	res, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if res {
		return utils_error.ErrMessage(response.FAILD, "该等级以存在")
	}

	return
}
