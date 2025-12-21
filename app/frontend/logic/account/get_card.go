package account

import (
	"context"
	dao_account "server/app/frontend/dao/account"
)

// GetCard implements service.IAccount.
func (s *sAccount) GetCard(ctx context.Context, id int64) (res *dao_account.CardList, err error) {
	// err = dao.SysWitkeyCard.Ctx(ctx).Where(dao.SysWitkeyCard.Columns().
	// 	Id, id).Scan(&res)
	return
}
