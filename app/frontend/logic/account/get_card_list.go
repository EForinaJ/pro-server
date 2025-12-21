package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_account "server/app/frontend/dao/account"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetCardList implements service.IAccount.
func (s *sAccount) GetCardList(ctx context.Context) (res []*dao_account.CardList, err error) {
	m := dao.SysUserCard.Ctx(ctx).
		Where(dao.SysUserCard.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysUserCard.Columns().CreateTime)

	var list []*entity.SysUserCard
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = make([]*dao_account.CardList, len(list))

	for i, v := range list {
		var entity *dao_account.CardList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		name := gconv.String(v.Name)
		entity.Name = gstr.SubStr(name, 0, 3) + "**"

		res[i] = entity
	}

	return
}
