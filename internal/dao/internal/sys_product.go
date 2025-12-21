// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductDao is the data access object for the table sys_product.
type SysProductDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysProductColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductColumns defines and stores column names for the table sys_product.
type SysProductColumns struct {
	Id          string //
	Pic         string //
	Type        string // 1护航，2代肝
	CategoryId  string //
	GameId      string //
	Name        string //
	Description string //
	Content     string //
	WitkeyCount string //
	Price       string //
	Unit        string //
	SalesCount  string //
	Rate        string //
	SkuType     string //
	Status      string // 状态：1-禁用，2-启用
	CreateTime  string //
	UpdateTime  string //
}

// sysProductColumns holds the columns for the table sys_product.
var sysProductColumns = SysProductColumns{
	Id:          "id",
	Pic:         "pic",
	Type:        "type",
	CategoryId:  "category_id",
	GameId:      "game_id",
	Name:        "name",
	Description: "description",
	Content:     "content",
	WitkeyCount: "witkey_count",
	Price:       "price",
	Unit:        "unit",
	SalesCount:  "sales_count",
	Rate:        "rate",
	SkuType:     "sku_type",
	Status:      "status",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewSysProductDao creates and returns a new DAO object for table data access.
func NewSysProductDao() *SysProductDao {
	return &SysProductDao{
		group:   "default",
		table:   "sys_product",
		columns: sysProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysProductDao) Columns() SysProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
