package dao_category

type List struct {
	Id          int64  `json:"id"     dc:"编号"`
	Name        string `json:"name" dc:"名称"`
	Slug        string `json:"slug" dc:"别名"`
	Cover       string `json:"cover"  dc:"封面"`
	Description string `json:"description" dc:"描述"`
}
