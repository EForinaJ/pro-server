package dto_order

type Create struct {
	GoodsId     int64  `p:"goodsId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"下单用户ID"`
	UserId      int64  `p:"userId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"下单商品ID"`
	Number      int    `p:"number" v:"required#请输入下单数量" dc:"下单数量"`
	Description string `p:"description" dc:"下单备注"`
}
