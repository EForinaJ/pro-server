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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IUser.
func (s *sUser) Create(ctx context.Context, req *dto_user.Create) (err error) {
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
	if req.Password == "" {
		req.Password = "123456"
	}
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	entity.Salt = newSalt
	entity.Password = newToken
	if req.Name == "" {
		req.Name = "新用户" + grand.S(6)
	}
	// 转换为 gtime.Time
	if req.Birthday == 0 {
		req.Birthday = 631123200000
	}
	birthday := gtime.NewFromTimeStamp(req.Birthday / 1000).UTC()
	entity.Birthday = birthday
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	//  获取配置
	userInfo, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.UserSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	userJosn := gjson.New(userInfo)
	//  获取默认头像和封面
	if req.Avatar == "" {
		entity.Avatar = userJosn.Get("avatar")
	}

	entity.Cover = userJosn.Get("cover")
	entity.Status = req.Status

	entity.LevelId = userJosn.Get("levelId").Int64()
	exp, err := tx.Model(dao.SysLevel.Table()).
		Where(dao.SysLevel.Columns().Id, userJosn.Get("levelId").Int64()).
		Value(dao.SysLevel.Columns().Experience)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.Experience = exp.Int()

	_, err = tx.Model(dao.SysUser.Table()).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
