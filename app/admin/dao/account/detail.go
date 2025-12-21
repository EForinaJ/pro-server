package dao_account

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id        int16       `json:"id" dc:"账户id"`
	Name      string      `json:"name" dc:"账户昵称"`
	Gender    string      `json:"gender" dc:"账户性别"`
	Avatar    string      `json:"avatar" dc:"账户头像"`
	LoginTime *gtime.Time `json:"loginTime" dc:"登录时间"`
	LoginIp   string      `json:"LoginIp" dc:"登录地址"`
	// Permission []string    `json:"permission" dc:"接口权限"`
}
