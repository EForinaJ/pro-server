// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRechargeDao is the data access object for the table sys_recharge.
type SysRechargeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysRechargeColumns // columns contains all the column names of Table for convenient usage.
}

// SysRechargeColumns defines and stores column names for the table sys_recharge.
type SysRechargeColumns struct {
	Id         string //
	UserId     string //
	Code       string //
	Amount     string //
	PayMode    string //
	Status     string //
	CreateTime string //
	UpdateTime string //
	Remark     string //
}

// sysRechargeColumns holds the columns for the table sys_recharge.
var sysRechargeColumns = SysRechargeColumns{
	Id:         "id",
	UserId:     "user_id",
	Code:       "code",
	Amount:     "amount",
	PayMode:    "pay_mode",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewSysRechargeDao creates and returns a new DAO object for table data access.
func NewSysRechargeDao() *SysRechargeDao {
	return &SysRechargeDao{
		group:   "default",
		table:   "sys_recharge",
		columns: sysRechargeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysRechargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysRechargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysRechargeDao) Columns() SysRechargeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysRechargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysRechargeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysRechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
