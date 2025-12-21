package witkey

import (
	"context"
	dao_witkey "server/app/admin/dao/witkey"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IWitkey.
func (s *sWitkey) GetEdit(ctx context.Context, id int64) (res *dao_witkey.Edit, err error) {
	info, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var detail *dao_witkey.Edit
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}
	return detail, nil
}
