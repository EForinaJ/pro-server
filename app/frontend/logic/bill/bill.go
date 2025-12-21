package bill

import (
	"server/app/frontend/service"
)

type sBill struct {
}

func init() {
	service.RegisterBill(&sBill{})
}
