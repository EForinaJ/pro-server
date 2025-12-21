package service

import (
	"context"
	dao_level "server/app/admin/dao/level"
	dto_level "server/app/admin/dto/level"
)

// 定义显示接口
type ILevel interface {
	// 获取所有
	GetAll(ctx context.Context) (res []*dao_level.List, err error)

	GetList(ctx context.Context, req *dto_level.Query) (total int, res []*dao_level.List, err error)
	Create(ctx context.Context, req *dto_level.Create) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_level.Edit, err error)

	Edit(ctx context.Context, req *dto_level.Edit) (err error)
	Delete(ctx context.Context, ids []int64) (err error)

	CheckCreate(ctx context.Context, req *dto_level.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_level.Edit) (err error)
}

// 定义接口变量
var localLevel ILevel

// 定义一个获取接口的方法
func Level() ILevel {
	if localLevel == nil {
		panic("implement not found for interface ILevel, forgot register?")
	}
	return localLevel
}

// 定义一个接口实现的注册方法
func RegisterLevel(i ILevel) {
	localLevel = i
}
