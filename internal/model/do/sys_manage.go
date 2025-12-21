// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysManage is the golang structure of table sys_manage for DAO operations like Where/Data.
type SysManage struct {
	g.Meta     `orm:"table:sys_manage, do:true"`
	Id         interface{} // 用户ID
	Name       interface{} // 用户昵称
	Phone      interface{} // 手机号码
	Sex        interface{} //
	Avatar     interface{} // 头像地址
	Password   interface{} // 密码
	Salt       interface{} // 密码盐
	Status     interface{} // 帐号状态（1停用,2正常）
	Qq         interface{} //
	Wechat     interface{} //
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Remark     interface{} // 备注
	DeleteTime *gtime.Time // 软删除
}
