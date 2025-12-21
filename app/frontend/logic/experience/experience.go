package experience

import "server/app/frontend/service"

type sExperience struct{}

func init() {
	service.RegisterExperience(&sExperience{})
}
