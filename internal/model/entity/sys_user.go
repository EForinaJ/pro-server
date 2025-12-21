// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id          int64       `json:"id"          orm:"id"          description:"用户ID"`            // 用户ID
	LevelId     int64       `json:"levelId"     orm:"level_id"    description:""`                //
	WxUnionId   string      `json:"wxUnionId"   orm:"wx_union_id" description:""`                //
	WxOpenId    string      `json:"wxOpenId"    orm:"wx_open_id"  description:""`                //
	Name        string      `json:"name"        orm:"name"        description:"用户昵称"`            // 用户昵称
	Phone       string      `json:"phone"       orm:"phone"       description:"手机号码"`            // 手机号码
	Sex         int         `json:"sex"         orm:"sex"         description:"用户性别（1男 2女 3未知）"` // 用户性别（1男 2女 3未知）
	Address     string      `json:"address"     orm:"address"     description:""`                //
	Birthday    *gtime.Time `json:"birthday"    orm:"birthday"    description:""`                //
	Avatar      string      `json:"avatar"      orm:"avatar"      description:"头像地址"`            // 头像地址
	Password    string      `json:"password"    orm:"password"    description:"密码"`              // 密码
	Salt        string      `json:"salt"        orm:"salt"        description:"密码盐"`             // 密码盐
	Cover       string      `json:"cover"       orm:"cover"       description:""`                //
	Experience  int64       `json:"experience"  orm:"experience"  description:""`                //
	Balance     float64     `json:"balance"     orm:"balance"     description:"余额"`              // 余额
	Commission  float64     `json:"commission"  orm:"commission"  description:""`                //
	Deposit     float64     `json:"deposit"     orm:"deposit"     description:""`                //
	Description string      `json:"description" orm:"description" description:""`                //
	Status      int         `json:"status"      orm:"status"      description:"帐号状态（1停用,2正常）"`   // 帐号状态（1停用,2正常）
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"`            // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新时间"`            // 更新时间
	DeleteTime  *gtime.Time `json:"deleteTime"  orm:"delete_time" description:"软删除"`             // 软删除
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`              // 备注
}
