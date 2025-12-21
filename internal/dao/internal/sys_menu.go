// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for the table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for the table sys_menu.
type SysMenuColumns struct {
	Id            string // 菜单ID
	PId           string // 父菜单ID
	MenuType      string // 菜单类型（1目录 2菜单）
	Name          string // 请求地址
	Path          string //
	Label         string // 菜单名称
	Component     string // 组件地址
	Icon          string //
	IsEnable      string //
	Sort          string //
	IsMenu        string //
	KeepAlive     string //
	IsHide        string // 菜单状态（2显示 1隐藏）
	IsHideTab     string //
	Link          string // 跳转
	IsIframe      string //
	ShowBadge     string //
	ShowTextBadge string //
	FixedTab      string //
	ActivePath    string //
	IsFullPage    string //
	AuthName      string //
	AuthLabel     string //
	AuthIcon      string //
	AuthSort      string //
}

// sysMenuColumns holds the columns for the table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:            "id",
	PId:           "p_id",
	MenuType:      "menu_type",
	Name:          "name",
	Path:          "path",
	Label:         "label",
	Component:     "component",
	Icon:          "icon",
	IsEnable:      "is_enable",
	Sort:          "sort",
	IsMenu:        "is_menu",
	KeepAlive:     "keep_alive",
	IsHide:        "is_hide",
	IsHideTab:     "is_hide_tab",
	Link:          "link",
	IsIframe:      "is_iframe",
	ShowBadge:     "show_badge",
	ShowTextBadge: "show_text_badge",
	FixedTab:      "fixed_tab",
	ActivePath:    "active_path",
	IsFullPage:    "is_full_page",
	AuthName:      "auth_name",
	AuthLabel:     "auth_label",
	AuthIcon:      "auth_icon",
	AuthSort:      "auth_sort",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
