package dto_product

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Name   string `p:"name" dc:"名称"`
	GameId int64  `p:"gameId" dc:"所属游戏"`
	Status int    `p:"status" dc:"商品状态"`
}
