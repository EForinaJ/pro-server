// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProjectDao is the data access object for the table sys_project.
type SysProjectDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysProjectColumns // columns contains all the column names of Table for convenient usage.
}

// SysProjectColumns defines and stores column names for the table sys_project.
type SysProjectColumns struct {
	Id         string //
	Code       string //
	OrderId    string //
	WitkeyId   string //
	Images     string //
	Money      string //
	ServiceFee string //
	Commission string //
	Status     string //
	StartTime  string //
	FinishTime string //
	CreateTime string //
}

// sysProjectColumns holds the columns for the table sys_project.
var sysProjectColumns = SysProjectColumns{
	Id:         "id",
	Code:       "code",
	OrderId:    "order_id",
	WitkeyId:   "witkey_id",
	Images:     "images",
	Money:      "money",
	ServiceFee: "service_fee",
	Commission: "commission",
	Status:     "status",
	StartTime:  "start_time",
	FinishTime: "finish_time",
	CreateTime: "create_time",
}

// NewSysProjectDao creates and returns a new DAO object for table data access.
func NewSysProjectDao() *SysProjectDao {
	return &SysProjectDao{
		group:   "default",
		table:   "sys_project",
		columns: sysProjectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProjectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProjectDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProjectDao) Columns() SysProjectColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProjectDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProjectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProjectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
