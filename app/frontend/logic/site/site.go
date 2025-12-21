package site

import (
	"server/app/frontend/service"
)

type sSite struct{}

func init() {
	service.RegisterSite(&sSite{})
}
