package comment

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_comment "server/app/frontend/dto/comment"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IComment.
func (s *sComment) Create(ctx context.Context, req *dto_comment.Create) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var entity *do.SysComment
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.UserId = ctx.Value("userId")
	entity.Ip = g.RequestFromCtx(ctx).GetClientIp()

	productId, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, req.OrderId).
		Value(dao.SysOrder.Columns().ProductId)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	entity.ProductId = productId.Int64()
	entity.Status = consts.StatusApply
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = tx.Model(dao.SysComment.Table()).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
