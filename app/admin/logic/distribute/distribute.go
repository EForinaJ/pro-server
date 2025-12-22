package distribute

import (
	"server/app/admin/service"
)

type sDistribute struct{}

func init() {
	service.RegisterDistribute(&sDistribute{})
}
