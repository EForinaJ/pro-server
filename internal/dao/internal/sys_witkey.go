// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWitkeyDao is the data access object for the table sys_witkey.
type SysWitkeyDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns SysWitkeyColumns // columns contains all the column names of Table for convenient usage.
}

// SysWitkeyColumns defines and stores column names for the table sys_witkey.
type SysWitkeyColumns struct {
	Id          string // 用户ID
	UserId      string //
	TitleId     string //
	GameId      string //
	Album       string //
	Rate        string //
	Commission  string //
	Description string //
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// sysWitkeyColumns holds the columns for the table sys_witkey.
var sysWitkeyColumns = SysWitkeyColumns{
	Id:          "id",
	UserId:      "user_id",
	TitleId:     "title_id",
	GameId:      "game_id",
	Album:       "album",
	Rate:        "rate",
	Commission:  "commission",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewSysWitkeyDao creates and returns a new DAO object for table data access.
func NewSysWitkeyDao() *SysWitkeyDao {
	return &SysWitkeyDao{
		group:   "default",
		table:   "sys_witkey",
		columns: sysWitkeyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWitkeyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWitkeyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWitkeyDao) Columns() SysWitkeyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWitkeyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWitkeyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWitkeyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
