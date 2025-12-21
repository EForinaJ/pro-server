package dao_witkey

type Edit struct {
	Id      int64    `json:"id" dc:"ID"`
	UserId  int64    `json:"userId"  dc:"所属用户"`
	TitleId int64    `json:"titleId"  dc:"头衔勋章"`
	GameId  int64    `json:"gameId" dc:"游戏领域"`
	Album   []string `json:"album" dc:"相册"`
	Rate    int      `json:"rate" dc:"评分"`
}
