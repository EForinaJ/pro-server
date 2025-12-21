package service

import (
	"context"
	dao_site "server/app/frontend/dao/site"
)

// 定义显示接口
type ISite interface {
	GetInfo(ctx context.Context) (res *dao_site.Detail, err error)
	GetDictData(ctx context.Context, code string) (res []*dao_site.DictData, err error)
}

// 定义接口变量
var localSite ISite

// 定义一个获取接口的方法
func Site() ISite {
	if localSite == nil {
		panic("implement not found for interface ISite, forgot register?")
	}
	return localSite
}

// 定义一个接口实现的注册方法
func RegisterSite(i ISite) {
	localSite = i
}
