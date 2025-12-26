// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWithdrawDao is the data access object for the table sys_withdraw.
type SysWithdrawDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysWithdrawColumns // columns contains all the column names of Table for convenient usage.
}

// SysWithdrawColumns defines and stores column names for the table sys_withdraw.
type SysWithdrawColumns struct {
	Id            string //
	Code          string //
	ManageId      string //
	WitkeyId      string //
	Amount        string //
	SettledAmount string //
	ServiceFee    string //
	Type          string //
	Number        string //
	Name          string //
	Status        string //
	Reason        string //
	CreateTime    string //
	UpdateTime    string //
}

// sysWithdrawColumns holds the columns for the table sys_withdraw.
var sysWithdrawColumns = SysWithdrawColumns{
	Id:            "id",
	Code:          "code",
	ManageId:      "manage_id",
	WitkeyId:      "witkey_id",
	Amount:        "amount",
	SettledAmount: "settled_amount",
	ServiceFee:    "service_fee",
	Type:          "type",
	Number:        "number",
	Name:          "name",
	Status:        "status",
	Reason:        "reason",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewSysWithdrawDao creates and returns a new DAO object for table data access.
func NewSysWithdrawDao() *SysWithdrawDao {
	return &SysWithdrawDao{
		group:   "default",
		table:   "sys_withdraw",
		columns: sysWithdrawColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWithdrawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWithdrawDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWithdrawDao) Columns() SysWithdrawColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWithdrawDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWithdrawDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
