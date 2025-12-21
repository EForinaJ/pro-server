package service

import (
	"context"
	dao_project "server/app/admin/dao/project"
	dto_project "server/app/admin/dto/project"
)

// 定义显示接口
type IProject interface {
	// 检查是否可以结算
	CheckSettlement(ctx context.Context, req *dto_project.Settlement) (err error)
	Settlement(ctx context.Context, req *dto_project.Settlement) (err error)

	// 检查是否可以取消
	CheckCancel(ctx context.Context, req *dto_project.Cancel) (err error)
	Cancel(ctx context.Context, req *dto_project.Cancel) (err error)
	GetDetail(ctx context.Context, id int64) (res *dao_project.Detail, err error)
	GetList(ctx context.Context, req *dto_project.Query) (total int, res []*dao_project.List, err error)
	GetLogs(ctx context.Context, req *dto_project.Log) (total int, res []*dao_project.Log, err error)
	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localProject IProject

// 定义一个获取接口的方法
func Project() IProject {
	if localProject == nil {
		panic("implement not found for interface IProject, forgot register?")
	}
	return localProject
}

// 定义一个接口实现的注册方法
func RegisterProject(i IProject) {
	localProject = i
}
