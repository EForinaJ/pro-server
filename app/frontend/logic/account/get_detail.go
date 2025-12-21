package account

import (
	"context"
	"fmt"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_account "server/app/frontend/dao/account"
	"server/internal/dao"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAccount.
func (s *sAccount) GetDetail(ctx context.Context) (res *dao_account.Detail, err error) {
	userId := ctx.Value("userId")
	options, err := g.Redis().Get(ctx, consts.FrontendAccout+gconv.String(userId))
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}

	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR)
		}
		return
	}

	detail, err := dao.SysUser.Ctx(ctx).Fields(dao.SysUser.Columns().Id,
		dao.SysUser.Columns().Name,
		dao.SysUser.Columns().Id,
		dao.SysUser.Columns().Sex,
		dao.SysUser.Columns().Birthday,
		dao.SysUser.Columns().Address,
		dao.SysUser.Columns().Description,
		dao.SysUser.Columns().CreateTime,
		dao.SysUser.Columns().Balance,
		dao.SysUser.Columns().Phone,
		dao.SysUser.Columns().LevelId,
		dao.SysUser.Columns().Experience,
		dao.SysUser.Columns().Avatar).
		Where(dao.SysUser.Columns().Id, userId).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	err = gconv.Scan(detail.Map(), &res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res.Birthday = gtime.New(detail.Map()[dao.SysUser.Columns().Birthday]).TimestampMilli()

	phone := gconv.String(detail.Map()[dao.SysUser.Columns().Phone])
	if phone != "" {
		res.Phone = gstr.SubStr(phone, 0, 3) + "****" + gstr.SubStr(phone, 7, 4)
	}

	res.Roles = make([]string, 0)

	exist, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().UserId, userId).Exist()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		res.Roles = append(res.Roles, "witkey")
	}

	// 获取用户等级
	level, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Id, detail.GMap().Get(dao.SysUser.Columns().LevelId)).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	var levelDetail *dao_account.Level
	err = gconv.Scan(level.Map(), &levelDetail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res.Level = levelDetail

	now := time.Now()
	key := fmt.Sprintf("sign:%d:%s", ctx.Value("userId"), now.Format("200601")) // 例如: sign:123:202511
	offset := now.Day() - 1

	// 执行GETBIT命令
	bit, err := g.Redis().Do(ctx, "GETBIT", key, offset)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	res.IsSign = bit.Int() == 1

	err = g.Redis().SetEX(ctx, consts.FrontendAccout+gconv.String(userId), res, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR)
	}

	return
}
