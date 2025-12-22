package distribute

import (
	"context"
	dto_distribute "server/app/admin/dto/distribute"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Create implements service.IDistribute.
func (s *sDistribute) Create(ctx context.Context, req *dto_distribute.Create) (err error) {

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Code, req.Code).Value(dao.SysOrder.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	_, err = dao.SysDistribute.Ctx(ctx).
		Data(g.Map{
			dao.SysDistribute.Columns().WitkeyId:   req.WitkeyId,
			dao.SysDistribute.Columns().OrderId:    order.Int64(),
			dao.SysDistribute.Columns().IsCancel:   consts.Not,
			dao.SysDistribute.Columns().ManageId:   ctx.Value("userId"),
			dao.SysDistribute.Columns().CreateTime: gtime.Now(),
		}).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
