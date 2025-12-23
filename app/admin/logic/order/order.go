package order

import (
	"server/app/admin/service"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(&sOrder{})
}
