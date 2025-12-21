// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserLogDao is the data access object for the table sys_user_log.
type SysUserLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysUserLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserLogColumns defines and stores column names for the table sys_user_log.
type SysUserLogColumns struct {
	Id         string //
	UserId     string //
	ManageId   string //
	Mode       string //
	Content    string //
	CreateTime string //
}

// sysUserLogColumns holds the columns for the table sys_user_log.
var sysUserLogColumns = SysUserLogColumns{
	Id:         "id",
	UserId:     "user_id",
	ManageId:   "manage_id",
	Mode:       "mode",
	Content:    "content",
	CreateTime: "create_time",
}

// NewSysUserLogDao creates and returns a new DAO object for table data access.
func NewSysUserLogDao() *SysUserLogDao {
	return &SysUserLogDao{
		group:   "default",
		table:   "sys_user_log",
		columns: sysUserLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserLogDao) Columns() SysUserLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
