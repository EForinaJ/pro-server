// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysManageRoleDao is the data access object for the table sys_manage_role.
type SysManageRoleDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysManageRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysManageRoleColumns defines and stores column names for the table sys_manage_role.
type SysManageRoleColumns struct {
	ManageId string //
	RoleId   string //
}

// sysManageRoleColumns holds the columns for the table sys_manage_role.
var sysManageRoleColumns = SysManageRoleColumns{
	ManageId: "manage_id",
	RoleId:   "role_id",
}

// NewSysManageRoleDao creates and returns a new DAO object for table data access.
func NewSysManageRoleDao() *SysManageRoleDao {
	return &SysManageRoleDao{
		group:   "default",
		table:   "sys_manage_role",
		columns: sysManageRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysManageRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysManageRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysManageRoleDao) Columns() SysManageRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysManageRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysManageRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysManageRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
