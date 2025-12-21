// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCredentialDao is the data access object for the table sys_credential.
type SysCredentialDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysCredentialColumns // columns contains all the column names of Table for convenient usage.
}

// SysCredentialColumns defines and stores column names for the table sys_credential.
type SysCredentialColumns struct {
	Id                     string //
	UserId                 string //
	CredentialType         string //
	CredentialName         string //
	CredentialNumber       string //
	CredentialFrontImage   string //
	CredentialAgainstImage string //
	Status                 string //
	CreateTime             string //
	UpdateTime             string //
}

// sysCredentialColumns holds the columns for the table sys_credential.
var sysCredentialColumns = SysCredentialColumns{
	Id:                     "id",
	UserId:                 "user_id",
	CredentialType:         "credential_type",
	CredentialName:         "credential_name",
	CredentialNumber:       "credential_number",
	CredentialFrontImage:   "credential_front_image",
	CredentialAgainstImage: "credential_against_image",
	Status:                 "status",
	CreateTime:             "create_time",
	UpdateTime:             "update_time",
}

// NewSysCredentialDao creates and returns a new DAO object for table data access.
func NewSysCredentialDao() *SysCredentialDao {
	return &SysCredentialDao{
		group:   "default",
		table:   "sys_credential",
		columns: sysCredentialColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCredentialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCredentialDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCredentialDao) Columns() SysCredentialColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCredentialDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCredentialDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysCredentialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
