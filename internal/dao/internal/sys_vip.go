// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysVipDao is the data access object for the table sys_vip.
type SysVipDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns SysVipColumns // columns contains all the column names of Table for convenient usage.
}

// SysVipColumns defines and stores column names for the table sys_vip.
type SysVipColumns struct {
	Id         string // 会员ID
	Name       string // 会员名称
	Level      string // 会员等级
	Icon       string // 会员图标
	Price      string // 会员价格
	Duration   string // 有效期(天)
	Benefits   string // 会员权益说明
	Discount   string // 折扣率(0-1)
	PointsRate string // 积分倍率
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Remark     string // 备注
}

// sysVipColumns holds the columns for the table sys_vip.
var sysVipColumns = SysVipColumns{
	Id:         "id",
	Name:       "name",
	Level:      "level",
	Icon:       "icon",
	Price:      "price",
	Duration:   "duration",
	Benefits:   "benefits",
	Discount:   "discount",
	PointsRate: "points_rate",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewSysVipDao creates and returns a new DAO object for table data access.
func NewSysVipDao() *SysVipDao {
	return &SysVipDao{
		group:   "default",
		table:   "sys_vip",
		columns: sysVipColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysVipDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysVipDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysVipDao) Columns() SysVipColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysVipDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysVipDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysVipDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
