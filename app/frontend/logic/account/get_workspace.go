package account

import (
	"context"
	dao_account "server/app/frontend/dao/account"
)

// GetWorkspace implements service.IAccount.
func (s *sAccount) GetWorkspace(ctx context.Context) (res *dao_account.Workspace, err error) {
	return nil, nil
}
