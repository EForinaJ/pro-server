// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrderWitkeyDao is the data access object for the table sys_order_witkey.
type SysOrderWitkeyDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SysOrderWitkeyColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrderWitkeyColumns defines and stores column names for the table sys_order_witkey.
type SysOrderWitkeyColumns struct {
	OrderId    string //
	WitkeyId   string //
	IsReplaced string //
	Reason     string //
	CreateTime string //
}

// sysOrderWitkeyColumns holds the columns for the table sys_order_witkey.
var sysOrderWitkeyColumns = SysOrderWitkeyColumns{
	OrderId:    "order_id",
	WitkeyId:   "witkey_id",
	IsReplaced: "is_replaced",
	Reason:     "reason",
	CreateTime: "create_time",
}

// NewSysOrderWitkeyDao creates and returns a new DAO object for table data access.
func NewSysOrderWitkeyDao() *SysOrderWitkeyDao {
	return &SysOrderWitkeyDao{
		group:   "default",
		table:   "sys_order_witkey",
		columns: sysOrderWitkeyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOrderWitkeyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOrderWitkeyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOrderWitkeyDao) Columns() SysOrderWitkeyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOrderWitkeyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOrderWitkeyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOrderWitkeyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
