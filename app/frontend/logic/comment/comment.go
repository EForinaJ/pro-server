package comment

import (
	"server/app/frontend/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(&sComment{})
}
