package service

import (
	"context"
	dao_game "server/app/admin/dao/game"
	dto_game "server/app/admin/dto/game"
)

// 定义显示接口
type IGame interface {
	// 获取所有
	GetAll(ctx context.Context) (res []*dao_game.List, err error)

	GetList(ctx context.Context, req *dto_game.Query) (total int, res []*dao_game.List, err error)

	Create(ctx context.Context, req *dto_game.Create) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_game.Edit, err error)

	Edit(ctx context.Context, req *dto_game.Edit) (err error)

	Delete(ctx context.Context, ids []int64) (err error)

	CheckCreate(ctx context.Context, req *dto_game.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_game.Edit) (err error)
}

// 定义接口变量
var localGame IGame

// 定义一个获取接口的方法
func Game() IGame {
	if localGame == nil {
		panic("implement not found for interface IGame, forgot register?")
	}
	return localGame
}

// 定义一个接口实现的注册方法
func RegisterGame(i IGame) {
	localGame = i
}
