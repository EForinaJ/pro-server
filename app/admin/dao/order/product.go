package dao_order

type Product struct {
	Id   int64  `json:"id"     dc:"商品编号"`
	Name string `json:"name" dc:"商品名称"`
	Pic  string `json:"pic" dc:"商品图片"`
	// Specifications string  `json:"specifications" dc:"商品规格"`
	Quantity  int     `json:"quantity" dc:"数量"`
	UnitPrice float64 `json:"unitPrice" dc:"单价"`
	Unit      string  `json:"unit" dc:"单价"`
	Game      string  `json:"game" dc:"游戏领域"`
	Category  string  `json:"category" dc:"商品分类"`
}
