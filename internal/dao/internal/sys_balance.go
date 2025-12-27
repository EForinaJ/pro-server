// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysBalanceDao is the data access object for the table sys_balance.
type SysBalanceDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysBalanceColumns // columns contains all the column names of Table for convenient usage.
}

// SysBalanceColumns defines and stores column names for the table sys_balance.
type SysBalanceColumns struct {
	Id         string //
	UserId     string //
	Related    string //
	After      string //
	Before     string //
	Amount     string //
	Mode       string //
	Type       string //
	Remark     string //
	CreateTime string //
}

// sysBalanceColumns holds the columns for the table sys_balance.
var sysBalanceColumns = SysBalanceColumns{
	Id:         "id",
	UserId:     "user_id",
	Related:    "related",
	After:      "after",
	Before:     "before",
	Amount:     "amount",
	Mode:       "mode",
	Type:       "type",
	Remark:     "remark",
	CreateTime: "create_time",
}

// NewSysBalanceDao creates and returns a new DAO object for table data access.
func NewSysBalanceDao() *SysBalanceDao {
	return &SysBalanceDao{
		group:   "default",
		table:   "sys_balance",
		columns: sysBalanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysBalanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysBalanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysBalanceDao) Columns() SysBalanceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysBalanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysBalanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysBalanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
