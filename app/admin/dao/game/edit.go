package dao_game

type Edit struct {
	Id          int64  `json:"id"     dc:"分类编号"`
	Pic         string `json:"pic"  dc:"分类封面"`
	Name        string `json:"name" dc:"分类名称"`
	Keywords    string `json:"keywords" dc:"关键字"`
	ShortName   string `json:"shortName" dc:"短名"`
	Description string `json:"description" dc:"分类描述"`
}
