// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserCardDao is the data access object for the table sys_user_card.
type SysUserCardDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysUserCardColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserCardColumns defines and stores column names for the table sys_user_card.
type SysUserCardColumns struct {
	Id         string //
	UserId     string //
	Type       string //
	Name       string //
	Number     string //
	NickName   string //
	CreateTime string //
	UpdateTime string //
}

// sysUserCardColumns holds the columns for the table sys_user_card.
var sysUserCardColumns = SysUserCardColumns{
	Id:         "id",
	UserId:     "user_id",
	Type:       "type",
	Name:       "name",
	Number:     "number",
	NickName:   "nick_name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysUserCardDao creates and returns a new DAO object for table data access.
func NewSysUserCardDao() *SysUserCardDao {
	return &SysUserCardDao{
		group:   "default",
		table:   "sys_user_card",
		columns: sysUserCardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserCardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserCardDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserCardDao) Columns() SysUserCardColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserCardDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserCardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserCardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
