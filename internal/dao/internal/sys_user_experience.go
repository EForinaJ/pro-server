// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserExperienceDao is the data access object for the table sys_user_experience.
type SysUserExperienceDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns SysUserExperienceColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserExperienceColumns defines and stores column names for the table sys_user_experience.
type SysUserExperienceColumns struct {
	Id         string //
	UserId     string //
	Experience string //
	Type       string //
	CreateTime string //
}

// sysUserExperienceColumns holds the columns for the table sys_user_experience.
var sysUserExperienceColumns = SysUserExperienceColumns{
	Id:         "id",
	UserId:     "user_id",
	Experience: "experience",
	Type:       "type",
	CreateTime: "create_time",
}

// NewSysUserExperienceDao creates and returns a new DAO object for table data access.
func NewSysUserExperienceDao() *SysUserExperienceDao {
	return &SysUserExperienceDao{
		group:   "default",
		table:   "sys_user_experience",
		columns: sysUserExperienceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserExperienceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserExperienceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserExperienceDao) Columns() SysUserExperienceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserExperienceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserExperienceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserExperienceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
