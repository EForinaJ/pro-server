package service

import (
	"context"
	dao_category "server/app/admin/dao/category"
	dto_category "server/app/admin/dto/category"
)

// 定义显示接口
type ICategory interface {
	GetList(ctx context.Context, req *dto_category.Query) (total int, res []*dao_category.List, err error)
	GetEdit(ctx context.Context, id int64) (res *dao_category.Edit, err error)

	Edit(ctx context.Context, req *dto_category.Edit) (err error)
	Delete(ctx context.Context, ids []int64) (err error)
	Create(ctx context.Context, req *dto_category.Create) (err error)

	CheckCreate(ctx context.Context, req *dto_category.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_category.Edit) (err error)
}

// 定义接口变量
var localCategory ICategory

// 定义一个获取接口的方法
func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

// 定义一个接口实现的注册方法
func RegisterCategory(i ICategory) {
	localCategory = i
}
