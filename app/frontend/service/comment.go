package service

import (
	"context"
	dto_comment "server/app/frontend/dto/comment"
)

// 定义显示接口
type IComment interface {
	CheckCreate(ctx context.Context, req *dto_comment.Create) (err error)
	Create(ctx context.Context, req *dto_comment.Create) (err error)
}

// 定义接口变量
var localComment IComment

// 定义一个获取接口的方法
func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

// 定义一个接口实现的注册方法
func RegisterComment(i IComment) {
	localComment = i
}
