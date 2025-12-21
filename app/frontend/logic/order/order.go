package order

import (
	"server/app/frontend/service"
)

type sOrder struct {
}

func init() {
	service.RegisterOrder(&sOrder{})
}
