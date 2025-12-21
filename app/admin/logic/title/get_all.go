package title

import (
	"context"
	dao_title "server/app/admin/dao/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetAll implements service.ITitle.
func (s *sTitle) GetAll(ctx context.Context, gameId int64) (res []*dao_title.OptionsList, err error) {
	err = dao.SysTitle.Ctx(ctx).Where(dao.SysTitle.Columns().GameId, gameId).
		Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
