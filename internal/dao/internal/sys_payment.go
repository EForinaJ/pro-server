// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPaymentDao is the data access object for the table sys_payment.
type SysPaymentDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysPaymentColumns // columns contains all the column names of Table for convenient usage.
}

// SysPaymentColumns defines and stores column names for the table sys_payment.
type SysPaymentColumns struct {
	Id         string //
	RelatedId  string //
	Code       string //
	Type       string //
	Money      string //
	Mode       string //
	CreateTime string //
}

// sysPaymentColumns holds the columns for the table sys_payment.
var sysPaymentColumns = SysPaymentColumns{
	Id:         "id",
	RelatedId:  "related_id",
	Code:       "code",
	Type:       "type",
	Money:      "money",
	Mode:       "mode",
	CreateTime: "create_time",
}

// NewSysPaymentDao creates and returns a new DAO object for table data access.
func NewSysPaymentDao() *SysPaymentDao {
	return &SysPaymentDao{
		group:   "default",
		table:   "sys_payment",
		columns: sysPaymentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPaymentDao) Columns() SysPaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPaymentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysPaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
