package dao_aftersales

type Product struct {
	Id             int64   `json:"id"     dc:"商品编号"`
	Name           string  `json:"name" dc:"商品名称"`
	Pic            string  `json:"pic" dc:"商品图片"`
	Specifications string  `json:"specifications" dc:"商品规格"`
	Quantity       int     `json:"quantity" dc:"数量"`
	UnitPrice      float64 `json:"unitPrice" dc:"单价"`
	Game           string  `json:"game" dc:"游戏领域"`
	Type           string  `json:"type" dc:"商品类型"`
}
