package service

import (
	"context"
	dao_permission "server/app/admin/dao/permission"
	dto_permission "server/app/admin/dto/permission"
)

// 定义显示接口
type IPermission interface {
	// 添加关联
	AddRelated(ctx context.Context, id int64, permission []int64) (err error)
	// 获取关联下的Id
	GetRelatedId(ctx context.Context, id int64) (ids []int64, err error)
	// 获取所有
	GetAll(ctx context.Context) (res []*dao_permission.List, err error)
	// 添加内容
	Create(ctx context.Context, req *dto_permission.Create) (err error)
	// 获取所有
	GetEdit(ctx context.Context, id int64) (res *dao_permission.Edit, err error)
	// 修改内容
	Edit(ctx context.Context, req *dto_permission.Edit) (err error)
	// // 删除
	Delete(ctx context.Context, ids []int64) (err error)
	// 检查是否存在
	CheckCreate(ctx context.Context, req *dto_permission.Create) (err error)
	// // 排除ID当前外是否存在
	CheckEdit(ctx context.Context, req *dto_permission.Edit) (err error)
}

// 定义接口变量
var localPermission IPermission

// 定义一个获取接口的方法
func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IAccount, forgot register?")
	}
	return localPermission
}

// 定义一个接口实现的注册方法
func RegisterPermission(i IPermission) {
	localPermission = i
}
