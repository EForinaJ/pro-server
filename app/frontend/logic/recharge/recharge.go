package recharge

import (
	"server/app/frontend/service"
)

type sRecharge struct{}

func init() {
	service.RegisterRecharge(&sRecharge{})
}
