package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckChangeWitkey implements service.IOrder.
func (s *sOrder) CheckChangeWitkey(ctx context.Context, req *dto_order.ChangeWitkey) (err error) {

	oldWitkeyExist, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, req.OldId).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if !oldWitkeyExist {
		return utils_error.ErrMessage(response.FAILD, "旧威客不存在，未找到")
	}

	newWitkeyExist, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, req.NewId).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if !newWitkeyExist {
		return utils_error.ErrMessage(response.FAILD, "新威客不存在，未找到")
	}

	exist, err := dao.SysOrderWitkey.Ctx(ctx).
		Where(dao.SysOrderWitkey.Columns().OrderId, req.Id).
		Where(dao.SysOrderWitkey.Columns().WitkeyId, req.NewId).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "该订单以派发过该威客了")
	}

	return
}
