package dao_product

type Edit struct {
	Id            int64   `json:"id"     dc:"商品编号"`
	Name          string  `json:"name" dc:"商品名称"`
	GameId        int64   `json:"gameId" dc:"所属游戏id"`
	CategoryId    int64   `json:"categoryId" dc:"所属分类id"`
	Pic           string  `json:"pic" dc:"商品封面"`
	Description   string  `json:"description" dc:"商品描述"`
	Content       string  `json:"content" dc:"商品详情"`
	Price         float64 `json:"price" dc:"商品价格"`
	Unit          string  `json:"unit" dc:"服务单位"`
	Status        int     `json:"status" dc:"商品状态"`
	Rate          int     `json:"rate" dc:"评分"`
	Number        int     `json:"number"  dc:"服务人数"`
	WitkeyCount   int     `json:"witkeyCount"   dc:"服务人数"`
	SalesCount    int64   `json:"salesCount"    dc:"销售数量"`
	PurchaseLimit int     `json:"purchaseLimit"  dc:"限购次数"`
	// Specification []Specification `json:"specification" dc:"商品规格属性"`
	// Sku           []Sku           `json:"sku" dc:"规格明细"`
}
