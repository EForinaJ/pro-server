package user

import (
	"server/app/admin/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}
