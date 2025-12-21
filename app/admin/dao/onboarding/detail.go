package dao_onboarding

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id           int64       `json:"id" dc:"ID"`
	Manage       string      `json:"manage" dc:"管理员"`
	User         string      `json:"user" dc:"所属用户"`
	WxNumber     string      `json:"wxNumber" dc:"微信号"`
	Game         string      `json:"game" dc:"游戏领域"`
	Title        string      `json:"title" dc:"头衔勋章"`
	Images       []string    `json:"images" dc:"认证图片"`
	RefuseReason string      `json:"refuseReason" dc:"拒绝原因"`
	Status       int         `json:"status" dc:"帐号状态（1待审核,2通过,3未通过）"`
	CreateTime   *gtime.Time `json:"createTime" dc:"创建时间"`
}
