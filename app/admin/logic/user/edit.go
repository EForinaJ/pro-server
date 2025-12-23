package user

import (
	"context"
	dto_user "server/app/admin/dto/user"
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

// Edit implements service.IUser.
func (s *sUser) Edit(ctx context.Context, req *dto_user.Edit) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var entity *do.SysUser
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	if req.Password != "" {
		newSalt := grand.S(6)
		newToken := consts.SYSTEMNAME + req.Password + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity.Salt = newSalt
		entity.Password = newToken
	} else {
		password, err := tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id, req.Id).Value(dao.SysUser.Columns().Password)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity.Password = password.String()
	}

	// 转换为 gtime.Time
	if req.Birthday == 0 {
		req.Birthday = 631123200000
	}
	birthday := gtime.NewFromTimeStamp(req.Birthday / 1000).UTC()
	entity.Birthday = birthday

	// userLevel, err := tx.Model(dao.SysUser.Table()).
	// 	Where(dao.SysUser.Columns().Id, req.Id).
	// 	Value(dao.SysUser.Columns().LevelId)
	// if err != nil {
	// 	return utils_error.Err(response.DB_SAVE_ERROR)
	// }
	// if userLevel.Int64() != req.LevelId {
	// 	exp, err := tx.Model(dao.SysLevel.Table()).
	// 		Where(dao.SysLevel.Columns().Id, req.LevelId).
	// 		Value(dao.SysLevel.Columns().Experience)
	// 	if err != nil {
	// 		return utils_error.Err(response.DB_SAVE_ERROR)
	// 	}
	// 	entity.Experience = exp.Int()
	// }

	entity.UpdateTime = gtime.Now()
	_, err = tx.Model(dao.SysUser.Table()).
		WherePri(req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
