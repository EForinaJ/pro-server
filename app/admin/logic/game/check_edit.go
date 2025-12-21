package game

import (
	"context"
	dto_game "server/app/admin/dto/game"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckEdit implements service.IGame.
func (s *sGame) CheckEdit(ctx context.Context, req *dto_game.Edit) (err error) {
	res, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Name, req.Name).
		WhereNotIn(dao.SysGame.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if res {
		return utils_error.ErrMessage(response.FAILD, "游戏名称已存在")
	}
	return
}
