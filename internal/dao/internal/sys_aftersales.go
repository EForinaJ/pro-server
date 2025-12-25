// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAftersalesDao is the data access object for the table sys_aftersales.
type SysAftersalesDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysAftersalesColumns // columns contains all the column names of Table for convenient usage.
}

// SysAftersalesColumns defines and stores column names for the table sys_aftersales.
type SysAftersalesColumns struct {
	Id         string //
	Code       string //
	ManageId   string //
	OrderId    string //
	Type       string //
	Amount     string //
	Reason     string //
	Status     string //
	CreateTime string //
}

// sysAftersalesColumns holds the columns for the table sys_aftersales.
var sysAftersalesColumns = SysAftersalesColumns{
	Id:         "id",
	Code:       "code",
	ManageId:   "manage_id",
	OrderId:    "order_id",
	Type:       "type",
	Amount:     "amount",
	Reason:     "reason",
	Status:     "status",
	CreateTime: "create_time",
}

// NewSysAftersalesDao creates and returns a new DAO object for table data access.
func NewSysAftersalesDao() *SysAftersalesDao {
	return &SysAftersalesDao{
		group:   "default",
		table:   "sys_aftersales",
		columns: sysAftersalesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAftersalesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAftersalesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAftersalesDao) Columns() SysAftersalesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAftersalesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAftersalesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysAftersalesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
