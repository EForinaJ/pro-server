package service

import (
	"context"
	dao_title "server/app/admin/dao/title"
	dto_title "server/app/admin/dto/title"
)

// 定义显示接口
type ITitle interface {
	GetAll(ctx context.Context, categoryId int64) (res []*dao_title.OptionsList, err error)

	GetList(ctx context.Context, req *dto_title.Query) (total int, res []*dao_title.List, err error)

	Create(ctx context.Context, req *dto_title.Create) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_title.Edit, err error)

	Edit(ctx context.Context, req *dto_title.Edit) (err error)

	Delete(ctx context.Context, ids []int64) (err error)

	CheckCreate(ctx context.Context, req *dto_title.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_title.Edit) (err error)
}

// 定义接口变量
var localTitle ITitle

// 定义一个获取接口的方法
func Title() ITitle {
	if localTitle == nil {
		panic("implement not found for interface ITitle, forgot register?")
	}
	return localTitle
}

// 定义一个接口实现的注册方法
func RegisterTitle(i ITitle) {
	localTitle = i
}
