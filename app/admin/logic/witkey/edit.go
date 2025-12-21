package witkey

import (
	"context"
	dto_witkey "server/app/admin/dto/witkey"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Edit implements service.IWitkey.
func (s *sWitkey) Edit(ctx context.Context, req *dto_witkey.Edit) (err error) {
	_, err = dao.SysWitkey.Ctx(ctx).Data(g.Map{
		dao.SysWitkey.Columns().GameId:     req.GameId,
		dao.SysWitkey.Columns().TitleId:    req.TitleId,
		dao.SysWitkey.Columns().Rate:       req.Rate,
		dao.SysWitkey.Columns().UpdateTime: gtime.Now(),
		dao.SysWitkey.Columns().Album:      req.Album,
	}).Where(dao.SysWitkey.Columns().Id, req.Id).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
