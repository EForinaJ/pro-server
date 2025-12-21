package category

import (
	"context"
	dao_category "server/app/admin/dao/category"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetEdit implements service.ICategory.
func (s *sCategory) GetEdit(ctx context.Context, id int64) (res *dao_category.Edit, err error) {
	err = dao.SysCategory.Ctx(ctx).Where(dao.SysCategory.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
