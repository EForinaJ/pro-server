// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrderDao is the data access object for the table sys_order.
type SysOrderDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SysOrderColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrderColumns defines and stores column names for the table sys_order.
type SysOrderColumns struct {
	Id             string //
	Code           string //
	UserId         string //
	ProductId      string //
	Specifications string //
	WitkeyCount    string //
	Quantity       string //
	Unit           string //
	UnitPrice      string //
	TotalAmount    string // 订单总额
	DiscountAmount string // 优惠金额
	ActualAmount   string // 实付金额
	PayAmount      string //
	Requirements   string //
	Status         string //
	PayMode        string //
	PayStatus      string //
	PayTime        string //
	StartTime      string //
	FinishTime     string //
	CreateTime     string //
	UpdateTime     string //
	DeleteTime     string //
	Remark         string //
}

// sysOrderColumns holds the columns for the table sys_order.
var sysOrderColumns = SysOrderColumns{
	Id:             "id",
	Code:           "code",
	UserId:         "user_id",
	ProductId:      "product_id",
	Specifications: "specifications",
	WitkeyCount:    "witkey_count",
	Quantity:       "quantity",
	Unit:           "unit",
	UnitPrice:      "unit_price",
	TotalAmount:    "total_amount",
	DiscountAmount: "discount_amount",
	ActualAmount:   "actual_amount",
	PayAmount:      "pay_amount",
	Requirements:   "requirements",
	Status:         "status",
	PayMode:        "pay_mode",
	PayStatus:      "pay_status",
	PayTime:        "pay_time",
	StartTime:      "start_time",
	FinishTime:     "finish_time",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	DeleteTime:     "delete_time",
	Remark:         "remark",
}

// NewSysOrderDao creates and returns a new DAO object for table data access.
func NewSysOrderDao() *SysOrderDao {
	return &SysOrderDao{
		group:   "default",
		table:   "sys_order",
		columns: sysOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOrderDao) Columns() SysOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
