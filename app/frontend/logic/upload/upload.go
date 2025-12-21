package upload

import (
	"server/app/frontend/service"
)

type sUpload struct{}

func init() {
	service.RegisterUpload(&sUpload{})
}
