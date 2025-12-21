package service

import (
	"context"
	dao_category "server/app/frontend/dao/category"
)

// 定义显示接口
type ICategory interface {
	GetAll(ctx context.Context) (res []*dao_category.List, err error)
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
