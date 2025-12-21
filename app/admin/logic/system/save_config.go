package system

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SaveConfig implements service.ISystem.
func (s *sSystem) SaveConfig(ctx context.Context, key string, name string, value interface{}) (err error) {
	entites := do.SysConfig{
		Key:   key,
		Name:  name,
		Value: value,
	}
	isExist, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, key).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	// 判断配置项是否存在
	if isExist {
		entites.UpdateTime = gtime.Now()
		_, err := g.Redis().Del(ctx, consts.BaseSetting)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		_, err = g.Redis().Del(ctx, "site")
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		_, err = dao.SysConfig.Ctx(ctx).
			Where(dao.SysConfig.Columns().Key, key).
			Data(entites).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
	} else {
		entites.CreateTime = gtime.Now()
		_, err = dao.SysConfig.Ctx(ctx).
			Data(entites).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
	}
	return
}
