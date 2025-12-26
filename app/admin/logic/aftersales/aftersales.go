package aftersales

import (
	"server/app/admin/service"
)

type sAftersales struct{}

func init() {
	service.RegisterAftersales(&sAftersales{})
}
