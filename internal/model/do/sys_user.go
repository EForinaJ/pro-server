// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta      `orm:"table:sys_user, do:true"`
	Id          interface{} // 用户ID
	LevelId     interface{} //
	WxUnionId   interface{} //
	WxOpenId    interface{} //
	Name        interface{} // 用户昵称
	Phone       interface{} // 手机号码
	Sex         interface{} // 用户性别（1男 2女 3未知）
	Address     interface{} //
	Birthday    *gtime.Time //
	Avatar      interface{} // 头像地址
	Password    interface{} // 密码
	Salt        interface{} // 密码盐
	Cover       interface{} //
	Experience  interface{} //
	Balance     interface{} // 余额
	Commission  interface{} //
	Deposit     interface{} //
	Description interface{} //
	Status      interface{} // 帐号状态（1停用,2正常）
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	DeleteTime  *gtime.Time // 软删除
	Remark      interface{} // 备注
}
