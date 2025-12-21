package manage

import (
	"server/app/admin/service"
)

type sManage struct{}

func init() {
	service.RegisterManage(&sManage{})
}
