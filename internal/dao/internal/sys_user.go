// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	Id          string // 用户ID
	LevelId     string //
	WxUnionId   string //
	WxOpenId    string //
	Name        string // 用户昵称
	Phone       string // 手机号码
	Sex         string // 用户性别（1男 2女 3未知）
	Address     string //
	Birthday    string //
	Avatar      string // 头像地址
	Password    string // 密码
	Salt        string // 密码盐
	Cover       string //
	Experience  string //
	Balance     string // 余额
	Commission  string //
	Deposit     string //
	Description string //
	Status      string // 帐号状态（1停用,2正常）
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	DeleteTime  string // 软删除
	Remark      string // 备注
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	Id:          "id",
	LevelId:     "level_id",
	WxUnionId:   "wx_union_id",
	WxOpenId:    "wx_open_id",
	Name:        "name",
	Phone:       "phone",
	Sex:         "sex",
	Address:     "address",
	Birthday:    "birthday",
	Avatar:      "avatar",
	Password:    "password",
	Salt:        "salt",
	Cover:       "cover",
	Experience:  "experience",
	Balance:     "balance",
	Commission:  "commission",
	Deposit:     "deposit",
	Description: "description",
	Status:      "status",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	DeleteTime:  "delete_time",
	Remark:      "remark",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
