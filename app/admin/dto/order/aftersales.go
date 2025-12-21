package dto_order

type Aftersales struct {
	OrderId int64    `p:"orderId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Type    int      `p:"type" v:"required#请选择售后类型" dc:"售后类型"`
	Content string   `p:"content" v:"required#请输入售后原因" dc:"售后原因"`
	Money   float64  `p:"money"  v:"required#请输入售后金额" dc:"售后金额"`
	Images  []string `p:"images"  dc:"售后文件"`
	Mode    int      `p:"mode" v:"required|in:1,2,3#请选择退款方式|状态值只能在1或2或3"    dc:"退款方式"`
	// ReceiptNum   string   `p:"receiptNum" v:"required#请输入收款收据单号"    dc:"收款收据单号"`
	// ReceiptFiles []string `p:"receiptFile" v:"required|array#请上传收款收据文件"    dc:"收款收据文件"`
	// PayMode      int      `p:"payMode" v:"required|in:1,2#请输入状态|状态值只能在1或2"    dc:"收款方式"`
}
