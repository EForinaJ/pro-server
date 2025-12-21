// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCategoryDao is the data access object for the table sys_category.
type SysCategoryDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// SysCategoryColumns defines and stores column names for the table sys_category.
type SysCategoryColumns struct {
	Id          string //
	GameId      string //
	Pid         string // 顶级分类
	Name        string // 分类名称
	Pic         string // 分类背景图
	Sort        string // 分类排序
	Description string // 分类描述
	CreateTime  string // 创建日期
	UpdateTime  string // 更新日期
}

// sysCategoryColumns holds the columns for the table sys_category.
var sysCategoryColumns = SysCategoryColumns{
	Id:          "id",
	GameId:      "game_id",
	Pid:         "pid",
	Name:        "name",
	Pic:         "pic",
	Sort:        "sort",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewSysCategoryDao creates and returns a new DAO object for table data access.
func NewSysCategoryDao() *SysCategoryDao {
	return &SysCategoryDao{
		group:   "default",
		table:   "sys_category",
		columns: sysCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCategoryDao) Columns() SysCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
