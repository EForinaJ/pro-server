package dao_product

type Sku struct {
	Code          string   `json:"code" dc:"规格标识"`
	Attrs         []Attrs  `json:"atts" dc:"规格属性"`
	AttrNames     []string `json:"attrNames" dc:"属性名称"`
	Values        []string `json:"values" dc:"对应的值"`
	Price         float64  `json:"price" dc:"售价"`
	OriginalPrice float64  `json:"originalPrice" dc:"原价"`
	Stock         int      `json:"stock" dc:"库存"`
}

type Attrs struct {
	Attr  string `json:"attr" dc:"属性"`
	Value string `json:"values" dc:"对应的值"`
}
