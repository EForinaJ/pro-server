package level

import (
	"server/app/admin/service"
)

type sLevel struct{}

func init() {
	service.RegisterLevel(&sLevel{})
}
