package title

import (
	"context"
	dto_title "server/app/admin/dto/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.ITitle.
func (s *sTitle) Create(ctx context.Context, req *dto_title.Create) (err error) {
	var entity *do.SysTitle
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.Name = req.Name
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysTitle.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
