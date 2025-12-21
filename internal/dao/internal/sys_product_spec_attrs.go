// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductSpecAttrsDao is the data access object for the table sys_product_spec_attrs.
type SysProductSpecAttrsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns SysProductSpecAttrsColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductSpecAttrsColumns defines and stores column names for the table sys_product_spec_attrs.
type SysProductSpecAttrsColumns struct {
	Id         string //
	ProductId  string //
	Name       string //
	CreateTime string //
}

// sysProductSpecAttrsColumns holds the columns for the table sys_product_spec_attrs.
var sysProductSpecAttrsColumns = SysProductSpecAttrsColumns{
	Id:         "id",
	ProductId:  "product_id",
	Name:       "name",
	CreateTime: "create_time",
}

// NewSysProductSpecAttrsDao creates and returns a new DAO object for table data access.
func NewSysProductSpecAttrsDao() *SysProductSpecAttrsDao {
	return &SysProductSpecAttrsDao{
		group:   "default",
		table:   "sys_product_spec_attrs",
		columns: sysProductSpecAttrsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProductSpecAttrsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProductSpecAttrsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProductSpecAttrsDao) Columns() SysProductSpecAttrsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProductSpecAttrsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProductSpecAttrsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProductSpecAttrsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
