package dict_data

import (
	"context"
	dao_dict_data "server/app/admin/dao/dict_data"
	dto_dict_data "server/app/admin/dto/dict_data"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetList implements service.IDictData.
func (s *sDictData) GetList(ctx context.Context, req *dto_dict_data.Query) (total int, res []*dao_dict_data.List, err error) {
	m := dao.SysDictData.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysDictData.Columns().CreateTime).
		Where(dao.SysDictData.Columns().Code, req.Code)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	err = m.Scan(&res)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
