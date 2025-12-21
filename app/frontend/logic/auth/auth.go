package auth

import (
	"server/app/frontend/service"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(&sAuth{})
}
