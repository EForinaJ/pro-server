// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPunishDao is the data access object for the table sys_punish.
type SysPunishDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns SysPunishColumns // columns contains all the column names of Table for convenient usage.
}

// SysPunishColumns defines and stores column names for the table sys_punish.
type SysPunishColumns struct {
	Id         string //
	Code       string //
	WitkeyId   string //
	ManageId   string //
	Money      string //
	Content    string //
	CreateTime string //
	UpdateTime string //
}

// sysPunishColumns holds the columns for the table sys_punish.
var sysPunishColumns = SysPunishColumns{
	Id:         "id",
	Code:       "code",
	WitkeyId:   "witkey_id",
	ManageId:   "manage_id",
	Money:      "money",
	Content:    "content",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysPunishDao creates and returns a new DAO object for table data access.
func NewSysPunishDao() *SysPunishDao {
	return &SysPunishDao{
		group:   "default",
		table:   "sys_punish",
		columns: sysPunishColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPunishDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPunishDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPunishDao) Columns() SysPunishColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPunishDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPunishDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysPunishDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
