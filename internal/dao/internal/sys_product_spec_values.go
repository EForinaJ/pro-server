// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductSpecValuesDao is the data access object for the table sys_product_spec_values.
type SysProductSpecValuesDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns SysProductSpecValuesColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductSpecValuesColumns defines and stores column names for the table sys_product_spec_values.
type SysProductSpecValuesColumns struct {
	Id         string //
	SpecAttsId string //
	Value      string //
	Sort       string //
	CreateTime string //
}

// sysProductSpecValuesColumns holds the columns for the table sys_product_spec_values.
var sysProductSpecValuesColumns = SysProductSpecValuesColumns{
	Id:         "id",
	SpecAttsId: "spec_atts_id",
	Value:      "value",
	Sort:       "sort",
	CreateTime: "create_time",
}

// NewSysProductSpecValuesDao creates and returns a new DAO object for table data access.
func NewSysProductSpecValuesDao() *SysProductSpecValuesDao {
	return &SysProductSpecValuesDao{
		group:   "default",
		table:   "sys_product_spec_values",
		columns: sysProductSpecValuesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProductSpecValuesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProductSpecValuesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProductSpecValuesDao) Columns() SysProductSpecValuesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProductSpecValuesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProductSpecValuesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProductSpecValuesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
