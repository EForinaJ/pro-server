package dict

import (
	"server/app/frontend/service"
)

type sDict struct{}

func init() {
	service.RegisterDict(&sDict{})
}
