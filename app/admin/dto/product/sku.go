package dto_product

type Sku struct {
	Code          string   `p:"code" dc:"规格标识"`
	Attrs         []Attrs  `p:"atts" dc:"规格属性"`
	AttrNames     []string `p:"attrNames" dc:"属性名称"`
	Values        []string `p:"values" dc:"对应的值"`
	Price         float64  `p:"price" dc:"售价"`
	OriginalPrice float64  `p:"originalPrice" dc:"原价"`
	Stock         int      `p:"stock" dc:"库存"`
}

type Attrs struct {
	Attr  string `p:"attr" dc:"属性"`
	Value string `p:"values" dc:"对应的值"`
}
