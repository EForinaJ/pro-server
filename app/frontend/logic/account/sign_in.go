package account

import (
	"context"
	"fmt"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// SignIn implements service.IAccount.
func (s *sAccount) SignIn(ctx context.Context) (res int, err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return 0, utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	//  获取配置
	userInfo, err := tx.Model(dao.SysConfig.Table()).
		Where(dao.SysConfig.Columns().Key, consts.UserSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return 0, utils_error.Err(response.DB_READ_ERROR)
	}
	userJosn := gjson.New(userInfo)
	checkInExperience := userJosn.Get("checkInExperience")

	user, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Fields(dao.SysUser.Columns().Experience, dao.SysUser.Columns().LevelId).One()
	if err != nil {
		return 0, utils_error.Err(response.DB_READ_ERROR)
	}
	newExperience := grand.Intn(checkInExperience.Int())
	newUserExperience := gconv.Int(user.GMap().Get(dao.SysUser.Columns().Experience)) + newExperience

	var levelList []entity.SysLevel
	err = dao.SysLevel.Ctx(ctx).Scan(&levelList)
	if err != nil {
		return 0, utils_error.Err(response.DB_READ_ERROR)
	}
	var largestLess entity.SysLevel
	found := false
	for _, v := range levelList {
		if v.Experience < newUserExperience {
			if !found || v.Experience > largestLess.Experience {
				largestLess = v
				found = true
			}
		}
	}
	userEntity := g.Map{
		dao.SysUser.Columns().Experience: newUserExperience,
	}
	if gconv.Int64(user.GMap().Get(dao.SysUser.Columns().LevelId)) != largestLess.Id {
		userEntity[dao.SysUser.Columns().LevelId] = largestLess.Id
	}

	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Data(userEntity).Update()
	if err != nil {
		return 0, utils_error.Err(response.DB_SAVE_ERROR)
	}

	experienceEntity := g.Map{
		dao.SysUserExperience.Columns().UserId:     ctx.Value("userId"),
		dao.SysUserExperience.Columns().Experience: newExperience,
		dao.SysUserExperience.Columns().Type:       consts.ExperienceTypeCheckIn,
		dao.SysUserExperience.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserExperience.Table()).
		Data(experienceEntity).Insert()
	if err != nil {
		return 0, utils_error.Err(response.DB_SAVE_ERROR)
	}

	// 生成Redis键
	now := time.Now()
	key := fmt.Sprintf("sign:%d:%s", ctx.Value("userId"), now.Format("200601")) // 例如: sign:123:202511

	// 计算偏移量（本月第几天，从0开始）
	offset := now.Day() - 1 // 注意：Redis位图偏移从0开始，1号对应offset 0

	// 使用GoRedis的SetBit方法设置签到位
	_, err = g.Redis().Do(ctx, "SETBIT", key, offset, 1)
	if err != nil {
		return 0, utils_error.Err(response.CACHE_SAVE_ERROR)
	}

	// 可选：设置键的过期时间，避免无限增长（例如一个月后过期）
	expireTime := now.AddDate(0, 1, 0).Sub(now) // 大约一个月后过期
	_, err = g.Redis().Expire(ctx, key, int64(expireTime.Seconds()))
	if err != nil {
		return 0, utils_error.Err(response.CACHE_SAVE_ERROR)
	}
	return
}
