package service

import (
	"context"
	dao_experience "server/app/frontend/dao/experience"
	dto_experience "server/app/frontend/dto/experience"
)

// 定义显示接口
type IExperience interface {
	GetList(ctx context.Context, req *dto_experience.Query) (total int, res []*dao_experience.List, err error)
}

// 定义接口变量
var localExperience IExperience

// 定义一个获取接口的方法
func Experience() IExperience {
	if localExperience == nil {
		panic("implement not found for interface IExperience, forgot register?")
	}
	return localExperience
}

// 定义一个接口实现的注册方法
func RegisterExperience(i IExperience) {
	localExperience = i
}
