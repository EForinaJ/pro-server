package dict

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_dict "server/app/frontend/dao/dict"
	"server/internal/dao"
)

// GetAll implements service.IDict.
func (s *sDict) GetAll(ctx context.Context, code string) (res []*dao_dict.List, err error) {
	err = dao.SysDictData.
		Ctx(ctx).Where(dao.SysDictData.Columns().Code, code).
		Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
