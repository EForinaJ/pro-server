package settlement

import (
	"context"
	dao_settlement "server/app/admin/dao/settlement"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.ISettlement.
func (s *sSettlement) GetDetail(ctx context.Context, id int64) (res *dao_settlement.Detail, err error) {
	obj, err := dao.SysSettlement.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}
	// 2. 使用结构体转换代替手动映射
	var detail *dao_settlement.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.FAILD)
	}

	//  订单编号
	orderCode, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().OrderId)).
		Value(dao.SysOrder.Columns().Code)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Code = orderCode.String()

	//  威客
	witkey, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId)).
		Fields(dao.SysWitkey.Columns().UserId, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	user, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id,
			witkey.GMap().Get(dao.SysWitkey.Columns().UserId)).
		Value(dao.SysUser.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Witkey = user.String()

	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id,
			obj.GMap().Get(dao.SysSettlement.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Manage = manage.String()

	res = detail
	return
}
