package account

import (
	"server/app/frontend/service"
)

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}
