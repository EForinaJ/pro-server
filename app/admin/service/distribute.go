package service

import (
	"context"
	dao_distribute "server/app/admin/dao/distribute"
	dto_distribute "server/app/admin/dto/distribute"
)

// 定义显示接口
type IDistribute interface {
	GetList(ctx context.Context, req *dto_distribute.Query) (total int, res []*dao_distribute.List, err error)
	// GetEdit(ctx context.Context, id int64) (res *dao_distribute.Edit, err error)

	Create(ctx context.Context, req *dto_distribute.Create) (err error)
	// Edit(ctx context.Context, req *dto_distribute.Edit) (err error)
	// Delete(ctx context.Context, ids []int64) (err error)

	CheckCreate(ctx context.Context, req *dto_distribute.Create) (err error)
	// CheckEdit(ctx context.Context, req *dto_distribute.Edit) (err error)
}

// 定义接口变量
var localDistribute IDistribute

// 定义一个获取接口的方法
func Distribute() IDistribute {
	if localDistribute == nil {
		panic("implement not found for interface IDistribute, forgot register?")
	}
	return localDistribute
}

// 定义一个接口实现的注册方法
func RegisterDistribute(i IDistribute) {
	localDistribute = i
}
