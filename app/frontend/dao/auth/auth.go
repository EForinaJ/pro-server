package dao_auth

type LoginRes struct {
	Token  string `json:"token" dc:"登录Token"`
	Expire int    `json:"expire" dc:"过期时间"`
}
