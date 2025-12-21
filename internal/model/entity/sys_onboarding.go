// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOnboarding is the golang structure for table sys_onboarding.
type SysOnboarding struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	ManageId     int64       `json:"manageId"     orm:"manage_id"     description:""` //
	UserId       int64       `json:"userId"       orm:"user_id"       description:""` //
	TitleId      int64       `json:"titleId"      orm:"title_id"      description:""` //
	GameId       int64       `json:"gameId"       orm:"game_id"       description:""` //
	Images       string      `json:"images"       orm:"images"        description:""` //
	WxNumber     string      `json:"wxNumber"     orm:"wx_number"     description:""` //
	Status       int         `json:"status"       orm:"status"        description:""` //
	RefuseReason string      `json:"refuseReason" orm:"refuse_reason" description:""` //
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   description:""` //
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"   description:""` //
}
