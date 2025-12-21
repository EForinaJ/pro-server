package category

import (
	"context"
	dto_category "server/app/admin/dto/category"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCreate implements service.ICategory.
func (s *sCategory) CheckCreate(ctx context.Context, req *dto_category.Create) (err error) {
	res, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if res {
		return utils_error.ErrMessage(response.FAILD, "该分类以存在")
	}
	return
}
