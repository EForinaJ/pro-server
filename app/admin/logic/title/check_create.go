package title

import (
	"context"
	dto_title "server/app/admin/dto/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCreate implements service.ITitle.
func (s *sTitle) CheckCreate(ctx context.Context, req *dto_title.Create) (err error) {
	res, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().GameId, req.GameId).
		Where(dao.SysTitle.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if res {
		return utils_error.ErrMessage(response.FAILD, "头衔已存在")
	}
	return
}
