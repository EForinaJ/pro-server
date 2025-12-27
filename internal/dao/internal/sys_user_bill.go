// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserBillDao is the data access object for the table sys_user_bill.
type SysUserBillDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysUserBillColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserBillColumns defines and stores column names for the table sys_user_bill.
type SysUserBillColumns struct {
	Id         string //
	UserId     string //
	RelatedId  string //
	Code       string //
	Type       string //
	Amount     string //
	Mode       string //
	CreateTime string //
}

// sysUserBillColumns holds the columns for the table sys_user_bill.
var sysUserBillColumns = SysUserBillColumns{
	Id:         "id",
	UserId:     "user_id",
	RelatedId:  "related_id",
	Code:       "code",
	Type:       "type",
	Amount:     "amount",
	Mode:       "mode",
	CreateTime: "create_time",
}

// NewSysUserBillDao creates and returns a new DAO object for table data access.
func NewSysUserBillDao() *SysUserBillDao {
	return &SysUserBillDao{
		group:   "default",
		table:   "sys_user_bill",
		columns: sysUserBillColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserBillDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserBillDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserBillDao) Columns() SysUserBillColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserBillDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserBillDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserBillDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
