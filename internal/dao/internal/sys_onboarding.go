// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOnboardingDao is the data access object for the table sys_onboarding.
type SysOnboardingDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysOnboardingColumns // columns contains all the column names of Table for convenient usage.
}

// SysOnboardingColumns defines and stores column names for the table sys_onboarding.
type SysOnboardingColumns struct {
	Id           string //
	ManageId     string //
	UserId       string //
	TitleId      string //
	GameId       string //
	Images       string //
	WxNumber     string //
	Status       string //
	RefuseReason string //
	CreateTime   string //
	UpdateTime   string //
}

// sysOnboardingColumns holds the columns for the table sys_onboarding.
var sysOnboardingColumns = SysOnboardingColumns{
	Id:           "id",
	ManageId:     "manage_id",
	UserId:       "user_id",
	TitleId:      "title_id",
	GameId:       "game_id",
	Images:       "images",
	WxNumber:     "wx_number",
	Status:       "status",
	RefuseReason: "refuse_reason",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewSysOnboardingDao creates and returns a new DAO object for table data access.
func NewSysOnboardingDao() *SysOnboardingDao {
	return &SysOnboardingDao{
		group:   "default",
		table:   "sys_onboarding",
		columns: sysOnboardingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOnboardingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOnboardingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOnboardingDao) Columns() SysOnboardingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOnboardingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOnboardingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOnboardingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
