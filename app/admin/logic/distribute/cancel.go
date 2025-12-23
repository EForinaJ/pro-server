package distribute

import (
	"context"
	dto_distribute "server/app/admin/dto/distribute"
	"server/app/common/consts"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Cancel implements service.IDistribute.
func (s *sDistribute) Cancel(ctx context.Context, req *dto_distribute.Cancel) (err error) {
	_, err = dao.SysDistribute.Ctx(ctx).Where(dao.SysDistribute.Columns().Id, req.Id).Data(g.Map{
		dao.SysDistribute.Columns().IsCancel: consts.Yes,
		dao.SysDistribute.Columns().Reason:   req.Reason,
	}).Update()
	return
}
