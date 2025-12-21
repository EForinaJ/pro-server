package service

import (
	"context"
	dao_onboarding "server/app/admin/dao/onboarding"
	dto_onboarding "server/app/admin/dto/onboarding"
)

// 定义显示接口
type IOnboarding interface {
	GetDetail(ctx context.Context, id int64) (res *dao_onboarding.Detail, err error)
	GetList(ctx context.Context, req *dto_onboarding.Query) (total int, res []*dao_onboarding.List, err error)

	Apply(ctx context.Context, req *dto_onboarding.Apply) (err error)
	CheckExistWitkey(ctx context.Context, id int64) (res bool, err error)
}

// 定义接口变量
var localOnboarding IOnboarding

// 定义一个获取接口的方法
func Onboarding() IOnboarding {
	if localOnboarding == nil {
		panic("implement not found for interface IOnboarding, forgot register?")
	}
	return localOnboarding
}

// 定义一个接口实现的注册方法
func RegisterOnboarding(i IOnboarding) {
	localOnboarding = i
}
