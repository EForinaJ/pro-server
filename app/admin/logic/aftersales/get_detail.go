package aftersales

import (
	"context"
	dao_aftersales "server/app/admin/dao/aftersales"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAftersales.
func (s *sAftersales) GetDetail(ctx context.Context, id int64) (res *dao_aftersales.Detail, err error) {
	obj, err := dao.SysAftersales.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}

	// 2. 使用结构体转换代替手动映射
	var detail *dao_aftersales.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.FAILD)
	}

	order, err := service.Order().GetDetail(ctx, gconv.Int64(obj.GMap().Get(dao.SysAftersales.Columns().OrderId)))
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Order = order

	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id,
			obj.GMap().Get(dao.SysAftersales.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Manage = manage.String()

	return detail, nil
}
