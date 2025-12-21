package system

import (
	"context"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

type sSystem struct{}

// GetBase implements service.ISystem.
func (s *sSystem) GetOne(ctx context.Context, key string) (res interface{}, err error) {
	res, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, key).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

func init() {
	service.RegisterSystem(&sSystem{})
}
