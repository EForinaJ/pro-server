package dto_withdraw

type Apply struct {
	Id          int64    `p:"id"  v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"提现id"`
	ReceiptNum  string   `p:"receiptNum" v:"required#请输入回执收据单号"    dc:"回执收据单号"`
	ReceiptFile []string `p:"receiptFile" v:"required|array#请输入回执收据文件"    dc:"回执收据文件"`
	Status      int      `p:"status" v:"required|in:2,3#请输入状态|状态值只能在2或3"    dc:"提现状态"`
	Remark      string   `p:"remark"    dc:"备注"`
}
