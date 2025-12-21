package witkey

import (
	"server/app/admin/service"
)

type sWitkey struct{}

func init() {
	service.RegisterWitkey(&sWitkey{})
}
