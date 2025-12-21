package withdraw

import (
	"context"
	dao_withdraw "server/app/admin/dao/withdraw"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetInfo implements service.IWithdraw.
func (s *sWithdraw) GetInfo(ctx context.Context, id int64) (res *dao_withdraw.Detail, err error) {
	info, err := dao.SysWithdraw.Ctx(ctx).
		Where(dao.SysWithdraw.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}

	var entity *dao_withdraw.Detail
	err = gconv.Scan(info, &entity)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	receiptFiles := info.GMap().Get(dao.SysWithdraw.Columns().ReceiptFiles)
	if receiptFiles != nil {
		entity.ReceiptFiles = gconv.Strings(receiptFiles)
	} else {
		entity.ReceiptFiles = []string{}
	}

	// witkey, err := dao.SysWitkey.Ctx(ctx).
	// 	Where(dao.SysWitkey.Columns().Id, info.GMap().Get(dao.SysWithdraw.Columns().WitkeyId)).
	// 	Value(dao.SysWitkey.Columns().Name)
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// entity.Witkey = witkey.String()

	//
	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, info.GMap().Get(dao.SysWithdraw.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	entity.Manage = gconv.String(manage)

	return entity, nil
}
