package account

import (
	"context"
	dto_account "server/app/frontend/dto/account"
)

// EditCard implements service.IAccount.
func (s *sAccount) EditCard(ctx context.Context, req *dto_account.EditCard) (err error) {
	// _, err = dao.SysWitkeyCard.Ctx(ctx).
	// 	Where(dao.SysWitkeyCard.Columns().WitkeyId, ctx.Value("userId")).
	// 	Where(dao.SysWitkeyCard.Columns().Id, req.Id).
	// 	Data(g.Map{
	// 		dao.SysWitkeyCard.Columns().Name:       req.Name,
	// 		dao.SysWitkeyCard.Columns().Number:     req.Number,
	// 		dao.SysWitkeyCard.Columns().UpdateTime: gtime.Now(),
	// 	}).
	// 	Update()
	// if err != nil {
	// 	return utils_error.Err(response.DB_SAVE_ERROR)
	// }
	return
}
