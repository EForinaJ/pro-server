package account

import (
	"server/app/admin/service"
)

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}
