// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSettlementDao is the data access object for the table sys_settlement.
type SysSettlementDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// SysSettlementColumns defines and stores column names for the table sys_settlement.
type SysSettlementColumns struct {
	Id             string //
	Code           string //
	OrderId        string //
	WitkeyId       string //
	ManageId       string //
	Amount         string //
	Commission     string //
	ServiceCharge  string //
	Status         string //
	Reason         string //
	SettlementTime string //
	CreateTime     string //
	UpdateTime     string //
}

// sysSettlementColumns holds the columns for the table sys_settlement.
var sysSettlementColumns = SysSettlementColumns{
	Id:             "id",
	Code:           "code",
	OrderId:        "order_id",
	WitkeyId:       "witkey_id",
	ManageId:       "manage_id",
	Amount:         "amount",
	Commission:     "commission",
	ServiceCharge:  "service_charge",
	Status:         "status",
	Reason:         "reason",
	SettlementTime: "settlement_time",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewSysSettlementDao creates and returns a new DAO object for table data access.
func NewSysSettlementDao() *SysSettlementDao {
	return &SysSettlementDao{
		group:   "default",
		table:   "sys_settlement",
		columns: sysSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysSettlementDao) Columns() SysSettlementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
