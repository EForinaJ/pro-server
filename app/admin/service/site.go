package service

import (
	"context"
	dao_site "server/app/admin/dao/site"
)

// 定义显示接口
type ISite interface {
	GetInfo(ctx context.Context) (res *dao_site.Detail, err error)
	GetUserOptions(ctx context.Context, name string) (res []*dao_site.Options, err error)
	GetGameOptions(ctx context.Context) (res []*dao_site.Options, err error)
	GetCategoryOptions(ctx context.Context, id int64) (res []*dao_site.Options, err error)
	GetTitleOptions(ctx context.Context, id int64) (res []*dao_site.Options, err error)
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
