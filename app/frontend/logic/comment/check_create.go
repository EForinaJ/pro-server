package comment

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_comment "server/app/frontend/dto/comment"
	"server/internal/dao"
)

// CheckCreate implements service.IComment.
func (s *sComment) CheckCreate(ctx context.Context, req *dto_comment.Create) (err error) {
	exist, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, req.OrderId).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysOrder.Columns().Status, consts.OrderStatusComplete).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if !exist {
		return utils_error.ErrMessage(response.FAILD, "订单未完成")
	}

	exist, err = dao.SysComment.Ctx(ctx).
		Where(dao.SysComment.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysComment.Columns().OrderId, req.OrderId).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "订单已评价了")
	}

	return
}
