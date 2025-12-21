package recharge

import (
	"server/app/admin/service"
)

type sRecharge struct{}

func init() {
	service.RegisterRecharge(&sRecharge{})
}
