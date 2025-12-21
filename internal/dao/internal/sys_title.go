// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTitleDao is the data access object for the table sys_title.
type SysTitleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SysTitleColumns // columns contains all the column names of Table for convenient usage.
}

// SysTitleColumns defines and stores column names for the table sys_title.
type SysTitleColumns struct {
	Id             string //
	GameId         string //
	Name           string //
	ServicePercent string //
	Pic            string //
	Description    string //
	CreateTime     string //
	UpdateTime     string //
}

// sysTitleColumns holds the columns for the table sys_title.
var sysTitleColumns = SysTitleColumns{
	Id:             "id",
	GameId:         "game_id",
	Name:           "name",
	ServicePercent: "service_percent",
	Pic:            "pic",
	Description:    "description",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewSysTitleDao creates and returns a new DAO object for table data access.
func NewSysTitleDao() *SysTitleDao {
	return &SysTitleDao{
		group:   "default",
		table:   "sys_title",
		columns: sysTitleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTitleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTitleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTitleDao) Columns() SysTitleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTitleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTitleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysTitleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
