package level

import (
	"server/app/frontend/service"
)

type sLevel struct{}

func init() {
	service.RegisterLevel(&sLevel{})
}
