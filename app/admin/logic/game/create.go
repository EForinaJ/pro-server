package game

import (
	"context"
	dto_game "server/app/admin/dto/game"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IGame.
func (s *sGame) Create(ctx context.Context, req *dto_game.Create) (err error) {

	var entity *do.SysGame
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysGame.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
