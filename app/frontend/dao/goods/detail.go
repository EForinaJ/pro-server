package dao_goods

type Detail struct {
	Id          int64   `json:"id"     dc:"商品编号"`
	Name        string  `json:"name" dc:"商品名称"`
	Description string  `json:"description" dc:"商品描述"`
	Cover       string  `json:"cover"  dc:"商品封面"`
	Category    string  `json:"category" dc:"所属分类"`
	Type        string  `json:"type"  dc:"商品类型"`
	Price       float64 `json:"price"  dc:"商品价格"`
	Unit        string  `json:"unit" dc:"服务单位"`
	Content     string  `json:"content" dc:"服务内容"`
	SalesCount  int64   `json:"salesCount" dc:"派发数量"`
}
