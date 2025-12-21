package service

import (
	"context"
	dao_document "server/app/frontend/dao/document"
)

// 定义显示接口
type IDocument interface {
	// 登录接口
	GetDetail(ctx context.Context, id int64) (res *dao_document.Detail, err error)
}

// 定义接口变量
var localDocument IDocument

// 定义一个获取接口的方法
func Document() IDocument {
	if localDocument == nil {
		panic("implement not found for interface IDocument, forgot register?")
	}
	return localDocument
}

// 定义一个接口实现的注册方法
func RegisterDocument(i IDocument) {
	localDocument = i
}
