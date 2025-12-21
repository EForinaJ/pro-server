// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMediaDao is the data access object for the table sys_media.
type SysMediaDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SysMediaColumns // columns contains all the column names of Table for convenient usage.
}

// SysMediaColumns defines and stores column names for the table sys_media.
type SysMediaColumns struct {
	Id         string //
	Path       string // 存放路径
	Name       string // 文件名称
	OrName     string // 原始文件名称
	Size       string // 文件大小
	Ext        string // 文件后缀
	CreateTime string // 创建时间
	DeleteTime string //
	MediaType  string // 文件类型
}

// sysMediaColumns holds the columns for the table sys_media.
var sysMediaColumns = SysMediaColumns{
	Id:         "id",
	Path:       "path",
	Name:       "name",
	OrName:     "or_name",
	Size:       "size",
	Ext:        "ext",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
	MediaType:  "media_type",
}

// NewSysMediaDao creates and returns a new DAO object for table data access.
func NewSysMediaDao() *SysMediaDao {
	return &SysMediaDao{
		group:   "default",
		table:   "sys_media",
		columns: sysMediaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMediaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMediaDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMediaDao) Columns() SysMediaColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMediaDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMediaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysMediaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
