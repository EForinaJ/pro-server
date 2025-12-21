package category

import (
	"server/app/admin/service"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(&sCategory{})
}
