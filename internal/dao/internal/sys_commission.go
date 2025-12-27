// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCommissionDao is the data access object for the table sys_commission.
type SysCommissionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysCommissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysCommissionColumns defines and stores column names for the table sys_commission.
type SysCommissionColumns struct {
	Id         string //
	WitkeyId   string //
	Related    string //
	After      string //
	Before     string //
	Amount     string //
	Mode       string //
	Type       string //
	Remark     string //
	CreateTime string //
}

// sysCommissionColumns holds the columns for the table sys_commission.
var sysCommissionColumns = SysCommissionColumns{
	Id:         "id",
	WitkeyId:   "witkey_id",
	Related:    "related",
	After:      "after",
	Before:     "before",
	Amount:     "amount",
	Mode:       "mode",
	Type:       "type",
	Remark:     "remark",
	CreateTime: "create_time",
}

// NewSysCommissionDao creates and returns a new DAO object for table data access.
func NewSysCommissionDao() *SysCommissionDao {
	return &SysCommissionDao{
		group:   "default",
		table:   "sys_commission",
		columns: sysCommissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCommissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCommissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCommissionDao) Columns() SysCommissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCommissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCommissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysCommissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
