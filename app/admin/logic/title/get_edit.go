package title

import (
	"context"
	dao_title "server/app/admin/dao/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetEdit implements service.ITitle.
func (s *sTitle) GetEdit(ctx context.Context, id int64) (res *dao_title.Edit, err error) {
	err = dao.SysTitle.Ctx(ctx).Where(dao.SysTitle.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
