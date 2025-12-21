// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLevelDao is the data access object for the table sys_level.
type SysLevelDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SysLevelColumns // columns contains all the column names of Table for convenient usage.
}

// SysLevelColumns defines and stores column names for the table sys_level.
type SysLevelColumns struct {
	Id         string //
	Name       string //
	Level      string //
	Icon       string //
	Experience string //
	Benefits   string //
	CreateTime string //
	UpdateTime string //
}

// sysLevelColumns holds the columns for the table sys_level.
var sysLevelColumns = SysLevelColumns{
	Id:         "id",
	Name:       "name",
	Level:      "level",
	Icon:       "icon",
	Experience: "experience",
	Benefits:   "benefits",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysLevelDao creates and returns a new DAO object for table data access.
func NewSysLevelDao() *SysLevelDao {
	return &SysLevelDao{
		group:   "default",
		table:   "sys_level",
		columns: sysLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLevelDao) Columns() SysLevelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
