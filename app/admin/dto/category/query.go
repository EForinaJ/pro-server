package dto_category

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	GameId int64  `p:"gameId" dc:"游戏领域"`
	Name   string `p:"name" dc:"名称"`
}
