// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCategory is the golang structure for table sys_category.
type SysCategory struct {
	Id          int64       `json:"id"          orm:"id"          description:""`      //
	GameId      int64       `json:"gameId"      orm:"game_id"     description:""`      //
	Pid         int64       `json:"pid"         orm:"pid"         description:"顶级分类"`  // 顶级分类
	Name        string      `json:"name"        orm:"name"        description:"分类名称"`  // 分类名称
	Pic         string      `json:"pic"         orm:"pic"         description:"分类背景图"` // 分类背景图
	Sort        int         `json:"sort"        orm:"sort"        description:"分类排序"`  // 分类排序
	Description string      `json:"description" orm:"description" description:"分类描述"`  // 分类描述
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建日期"`  // 创建日期
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新日期"`  // 更新日期
}
