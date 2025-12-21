// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysManageDao is the data access object for the table sys_manage.
type SysManageDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns SysManageColumns // columns contains all the column names of Table for convenient usage.
}

// SysManageColumns defines and stores column names for the table sys_manage.
type SysManageColumns struct {
	Id         string // 用户ID
	Name       string // 用户昵称
	Phone      string // 手机号码
	Sex        string //
	Avatar     string // 头像地址
	Password   string // 密码
	Salt       string // 密码盐
	Status     string // 帐号状态（1停用,2正常）
	Qq         string //
	Wechat     string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Remark     string // 备注
	DeleteTime string // 软删除
}

// sysManageColumns holds the columns for the table sys_manage.
var sysManageColumns = SysManageColumns{
	Id:         "id",
	Name:       "name",
	Phone:      "phone",
	Sex:        "sex",
	Avatar:     "avatar",
	Password:   "password",
	Salt:       "salt",
	Status:     "status",
	Qq:         "qq",
	Wechat:     "wechat",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Remark:     "remark",
	DeleteTime: "delete_time",
}

// NewSysManageDao creates and returns a new DAO object for table data access.
func NewSysManageDao() *SysManageDao {
	return &SysManageDao{
		group:   "default",
		table:   "sys_manage",
		columns: sysManageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysManageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysManageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysManageDao) Columns() SysManageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysManageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysManageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysManageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
