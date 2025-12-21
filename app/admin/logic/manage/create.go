package manage

import (
	"context"
	dto_manage "server/app/admin/dto/manage"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IManage.
func (s *sManage) Create(ctx context.Context, req *dto_manage.Create) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var entity *do.SysManage
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	if req.Password == "" {
		req.Password = grand.S(6)
	}
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	entity.Salt = newSalt
	entity.Password = newToken

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	rs, err := tx.Model(dao.SysManage.Table()).
		Data(&entity).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	if len(req.Roles) > 0 {
		// 更新关联数据
		err = service.Role().AddRelated(ctx, tx, rid, req.Roles)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		//  更新casbin 权限内容
		_, err = tx.Model(dao.SysCasbin.Table()).
			Where(dao.SysCasbin.Columns().PType, "g").
			Where(dao.SysCasbin.Columns().V0, rid).Delete()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		roles, err := tx.Model(dao.SysRole.Table()).
			WhereIn(dao.SysRole.Columns().Id, req.Roles).
			Fields(dao.SysRole.Columns().Code).
			Array()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		entites := make([]g.Map, len(req.Roles))
		for i, v := range roles {
			entites[i] = g.Map{
				dao.SysCasbin.Columns().PType: "g",
				dao.SysCasbin.Columns().V0:    rid,
				dao.SysCasbin.Columns().V1:    v.String(),
			}
		}

		_, err = tx.Model(dao.SysCasbin.Table()).
			Data(entites).
			Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
	}

	return
}
