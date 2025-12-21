package project

import (
	"server/app/admin/service"
)

type sProject struct{}

func init() {
	service.RegisterProject(&sProject{})
}
