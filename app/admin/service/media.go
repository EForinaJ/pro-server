package service

import (
	"context"
	dao_media "server/app/admin/dao/media"
	dto_media "server/app/admin/dto/media"
)

// 定义显示接口
type IMedia interface {
	// 获取附件列表
	GetList(ctx context.Context, req *dto_media.Query) (total int, res []*dao_media.List, err error)
	// 删除附件
	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localMedia IMedia

// 定义一个获取接口的方法
func Media() IMedia {
	if localMedia == nil {
		panic("implement not found for interface IMedia, forgot register?")
	}
	return localMedia
}

// 定义一个接口实现的注册方法
func RegisterMedia(i IMedia) {
	localMedia = i
}
