package product

import (
	"server/app/admin/service"
)

type sProduct struct{}

func init() {
	service.RegisterProduct(&sProduct{})
}
