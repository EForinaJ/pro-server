package service

import (
	"context"
	dao_dict_type "server/app/admin/dao/dict_type"
	dto_dict_type "server/app/admin/dto/dict_type"
)

// 定义显示接口
type IDictType interface {
	// GetAll(ctx context.Context, categoryId int64) (res []*dao_title.List, err error)

	GetList(ctx context.Context, req *dto_dict_type.Query) (total int, res []*dao_dict_type.List, err error)

	Create(ctx context.Context, req *dto_dict_type.Create) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_dict_type.Edit, err error)

	Edit(ctx context.Context, req *dto_dict_type.Edit) (err error)

	Delete(ctx context.Context, ids []int64) (err error)

	// 检查是否存在
	CheckUnique(ctx context.Context, code string) (res bool, err error)

	// 排除ID是否存在
	CheckUniqueById(ctx context.Context, code string, id int64) (res bool, err error)
}

// 定义接口变量
var localDictType IDictType

// 定义一个获取接口的方法
func DictType() IDictType {
	if localDictType == nil {
		panic("implement not found for interface IDictType, forgot register?")
	}
	return localDictType
}

// 定义一个接口实现的注册方法
func RegisterDictType(i IDictType) {
	localDictType = i
}
