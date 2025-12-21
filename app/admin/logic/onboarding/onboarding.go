package onboarding

import (
	"server/app/admin/service"
)

type sOnboarding struct{}

func init() {
	service.RegisterOnboarding(&sOnboarding{})
}
