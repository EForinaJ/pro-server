// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrderLogDao is the data access object for the table sys_order_log.
type SysOrderLogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysOrderLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrderLogColumns defines and stores column names for the table sys_order_log.
type SysOrderLogColumns struct {
	Id         string //
	OrderId    string //
	ManageId   string //
	Type       string //
	Content    string //
	CreateTime string //
}

// sysOrderLogColumns holds the columns for the table sys_order_log.
var sysOrderLogColumns = SysOrderLogColumns{
	Id:         "id",
	OrderId:    "order_id",
	ManageId:   "manage_id",
	Type:       "type",
	Content:    "content",
	CreateTime: "create_time",
}

// NewSysOrderLogDao creates and returns a new DAO object for table data access.
func NewSysOrderLogDao() *SysOrderLogDao {
	return &SysOrderLogDao{
		group:   "default",
		table:   "sys_order_log",
		columns: sysOrderLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOrderLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOrderLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOrderLogDao) Columns() SysOrderLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOrderLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOrderLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOrderLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
