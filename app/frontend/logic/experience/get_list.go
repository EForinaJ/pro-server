package experience

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_experience "server/app/frontend/dao/experience"
	dto_experience "server/app/frontend/dto/experience"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IExperience.
func (s *sExperience) GetList(ctx context.Context, req *dto_experience.Query) (total int, res []*dao_experience.List, err error) {
	m := dao.SysUserExperience.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysUserExperience.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysUserExperience.Columns().CreateTime)

	if req.Type != 0 {
		m = m.Where(dao.SysUserExperience.Columns().Type, req.Type)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	var list []*entity.SysUserExperience
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = make([]*dao_experience.List, len(list))
	for i, v := range list {
		var entity *dao_experience.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}
	return
}
