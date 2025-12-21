package title

import (
	"server/app/admin/service"
)

type sTitle struct {
}

func init() {
	service.RegisterTitle(&sTitle{})
}
