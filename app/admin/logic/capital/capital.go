package capital

import (
	"server/app/admin/service"
)

type sCapital struct{}

func init() {
	service.RegisterCapital(&sCapital{})
}
