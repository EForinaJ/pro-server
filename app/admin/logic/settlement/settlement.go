package settlement

import (
	"server/app/admin/service"
)

type sSettlement struct{}

func init() {
	service.RegisterSettlement(&sSettlement{})
}
