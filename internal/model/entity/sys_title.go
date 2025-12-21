// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTitle is the golang structure for table sys_title.
type SysTitle struct {
	Id             int64       `json:"id"             orm:"id"              description:""` //
	GameId         int64       `json:"gameId"         orm:"game_id"         description:""` //
	Name           string      `json:"name"           orm:"name"            description:""` //
	ServicePercent int         `json:"servicePercent" orm:"service_percent" description:""` //
	Pic            string      `json:"pic"            orm:"pic"             description:""` //
	Description    string      `json:"description"    orm:"description"     description:""` //
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:""` //
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"     description:""` //
}
