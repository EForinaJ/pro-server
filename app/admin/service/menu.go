package service

import (
	"context"
	dao_menu "server/app/admin/dao/menu"
	dto_meun "server/app/admin/dto/meun"
)

// 定义显示接口
type IMenu interface {
	// 添加关联
	AddRelated(ctx context.Context, id int64, menus []int64) (err error)
	// 获取关联下的Id
	GetRelatedId(ctx context.Context, id int64) (ids []int64, err error)
	// 获取所有菜单
	GetAll(ctx context.Context) (res []*dao_menu.List, err error)
	// 创建菜单
	Create(ctx context.Context, req *dto_meun.Create) (err error)
	// 获取修改信息
	GetEdit(ctx context.Context, id int64) (res *dao_menu.Edit, err error)
	// 修改菜单
	Edit(ctx context.Context, req *dto_meun.Edit) (err error)
	// 删除菜单
	Delete(ctx context.Context, ids []int64) (err error)

	// 检查地址唯一
	CheckCreate(ctx context.Context, req *dto_meun.Create) (err error)
	// 检查特定菜单的地址唯一
	CheckEdit(ctx context.Context, req *dto_meun.Edit) (err error)
}

// 定义接口变量
var localMenu IMenu

// 定义一个获取接口的方法
func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

// 定义一个接口实现的注册方法
func RegisterMenu(i IMenu) {
	localMenu = i
}
