// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductSkuSpecRelationsDao is the data access object for the table sys_product_sku_spec_relations.
type SysProductSkuSpecRelationsDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of the current DAO.
	columns SysProductSkuSpecRelationsColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductSkuSpecRelationsColumns defines and stores column names for the table sys_product_sku_spec_relations.
type SysProductSkuSpecRelationsColumns struct {
	Id          string //
	SkuId       string //
	SpecAttrId  string //
	SpecValueId string //
	CreateTime  string //
}

// sysProductSkuSpecRelationsColumns holds the columns for the table sys_product_sku_spec_relations.
var sysProductSkuSpecRelationsColumns = SysProductSkuSpecRelationsColumns{
	Id:          "id",
	SkuId:       "sku_id",
	SpecAttrId:  "spec_attr_id",
	SpecValueId: "spec_value_id",
	CreateTime:  "create_time",
}

// NewSysProductSkuSpecRelationsDao creates and returns a new DAO object for table data access.
func NewSysProductSkuSpecRelationsDao() *SysProductSkuSpecRelationsDao {
	return &SysProductSkuSpecRelationsDao{
		group:   "default",
		table:   "sys_product_sku_spec_relations",
		columns: sysProductSkuSpecRelationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProductSkuSpecRelationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProductSkuSpecRelationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProductSkuSpecRelationsDao) Columns() SysProductSkuSpecRelationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProductSkuSpecRelationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProductSkuSpecRelationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProductSkuSpecRelationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
