package dao_category

type Edit struct {
	Id          int64  `json:"id"     dc:"分类编号"`
	GameId      int64  `json:"gameId" dc:"游戏领域"`
	Pic         string `json:"pic"  dc:"分类封面"`
	Name        string `json:"name" dc:"分类名称"`
	Description string `json:"description" dc:"分类描述"`
}
