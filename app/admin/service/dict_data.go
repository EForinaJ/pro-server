package service

import (
	"context"
	dao_dict_data "server/app/admin/dao/dict_data"
	dto_dict_data "server/app/admin/dto/dict_data"
)

// 定义显示接口
type IDictData interface {
	GetAll(ctx context.Context, code string) (res []*dao_dict_data.List, err error)

	GetList(ctx context.Context, req *dto_dict_data.Query) (total int, res []*dao_dict_data.List, err error)

	Create(ctx context.Context, req *dto_dict_data.Create) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_dict_data.Edit, err error)

	Edit(ctx context.Context, req *dto_dict_data.Edit) (err error)

	Delete(ctx context.Context, ids []int64) (err error)

	// 检查是否存在
	CheckUnique(ctx context.Context, code, key string) (res bool, err error)

	// 排除ID是否存在
	CheckUniqueById(ctx context.Context, code, key string, id int64) (res bool, err error)
}

// 定义接口变量
var localDictData IDictData

// 定义一个获取接口的方法
func DictData() IDictData {
	if localDictData == nil {
		panic("implement not found for interface IDictData, forgot register?")
	}
	return localDictData
}

// 定义一个接口实现的注册方法
func RegisterDictData(i IDictData) {
	localDictData = i
}
