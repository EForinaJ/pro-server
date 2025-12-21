package document

import "server/app/frontend/service"

type sDocument struct{}

func init() {
	service.RegisterDocument(&sDocument{})
}
