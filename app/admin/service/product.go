package service

import (
	"context"
	dao_product "server/app/admin/dao/product"
	dto_product "server/app/admin/dto/product"
)

// 定义显示接口
type IProduct interface {
	GetAllByName(ctx context.Context, name string) (res []*dao_product.List, err error)
	GetEdit(ctx context.Context, id int64) (res *dao_product.Edit, err error)
	GetList(ctx context.Context, req *dto_product.Query) (total int, res []*dao_product.List, err error)

	Create(ctx context.Context, req *dto_product.Create) (err error)
	Edit(ctx context.Context, req *dto_product.Edit) (err error)
	Delete(ctx context.Context, ids []int64) (err error)

	CheckSkuEmpty(ctx context.Context, req []dto_product.Sku) (err error)
	CheckSpecAttrEmpty(ctx context.Context, req []dto_product.Specification) (err error)
}

// 定义接口变量
var localProduct IProduct

// 定义一个获取接口的方法
func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

// 定义一个接口实现的注册方法
func RegisterProduct(i IProduct) {
	localProduct = i
}
