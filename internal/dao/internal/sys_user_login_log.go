// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserLoginLogDao is the data access object for the table sys_user_login_log.
type SysUserLoginLogDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SysUserLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserLoginLogColumns defines and stores column names for the table sys_user_login_log.
type SysUserLoginLogColumns struct {
	Id         string //
	UserId     string //
	Ip         string //
	CreateTime string //
}

// sysUserLoginLogColumns holds the columns for the table sys_user_login_log.
var sysUserLoginLogColumns = SysUserLoginLogColumns{
	Id:         "id",
	UserId:     "user_id",
	Ip:         "ip",
	CreateTime: "create_time",
}

// NewSysUserLoginLogDao creates and returns a new DAO object for table data access.
func NewSysUserLoginLogDao() *SysUserLoginLogDao {
	return &SysUserLoginLogDao{
		group:   "default",
		table:   "sys_user_login_log",
		columns: sysUserLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserLoginLogDao) Columns() SysUserLoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
