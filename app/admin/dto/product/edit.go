package dto_product

type Edit struct {
	Id            int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"商品id"`
	Name          string  `p:"name" v:"required#请输入商品名称" dc:"商品名称"`
	GameId        int64   `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属游戏id"`
	CategoryId    int64   `p:"categoryId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属分类id"`
	Pic           string  `p:"pic" v:"required#请上传商品封面" dc:"商品封面"`
	Description   string  `p:"description" dc:"商品描述"`
	Content       string  `p:"content" dc:"商品详情"`
	Price         float64 `p:"price" v:"required#请设置商品价格" dc:"商品价格"`
	Unit          string  `p:"unit" v:"required#请输入服务单位" dc:"服务单位"`
	Status        int     `p:"status" v:"required#请上设置商品状态" dc:"商品状态"`
	Rate          int     `p:"rate"     dc:"评分"`
	WitkeyCount   int     `p:"witkeyCount"    v:"required#请上设置服务人数"  dc:"服务人数"`
	SalesCount    int64   `p:"salesCount"    dc:"销售数量"`
	PurchaseLimit int     `p:"purchaseLimit"  dc:"限购次数"`
	// Specification []Specification `p:"specification" dc:"商品规格属性"`
	// Sku           []Sku           `p:"sku" dc:"规格明细"`
}
