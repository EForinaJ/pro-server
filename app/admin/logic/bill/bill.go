package bill

import (
	"server/app/admin/service"
)

type sBill struct{}

func init() {
	service.RegisterBill(&sBill{})
}
