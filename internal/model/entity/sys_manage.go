// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysManage is the golang structure for table sys_manage.
type SysManage struct {
	Id         int64       `json:"id"         orm:"id"          description:"用户ID"`          // 用户ID
	Name       string      `json:"name"       orm:"name"        description:"用户昵称"`          // 用户昵称
	Phone      string      `json:"phone"      orm:"phone"       description:"手机号码"`          // 手机号码
	Sex        int         `json:"sex"        orm:"sex"         description:""`              //
	Avatar     string      `json:"avatar"     orm:"avatar"      description:"头像地址"`          // 头像地址
	Password   string      `json:"password"   orm:"password"    description:"密码"`            // 密码
	Salt       string      `json:"salt"       orm:"salt"        description:"密码盐"`           // 密码盐
	Status     int         `json:"status"     orm:"status"      description:"帐号状态（1停用,2正常）"` // 帐号状态（1停用,2正常）
	Qq         string      `json:"qq"         orm:"qq"          description:""`              //
	Wechat     string      `json:"wechat"     orm:"wechat"      description:""`              //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`          // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`          // 更新时间
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`            // 备注
	DeleteTime *gtime.Time `json:"deleteTime" orm:"delete_time" description:"软删除"`           // 软删除
}
