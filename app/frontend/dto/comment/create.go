package dto_comment

type Create struct {
	OrderId int64    `p:"orderId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Rating  int      `p:"rating" v:"required|min:1|max:5|integer#请选择评分|评分最小为1|评分最大为5|评分必须是整型" dc:"评分"`
	Content string   `p:"content" dc:"评价内容内容"`
	Images  []string `p:"images" dc:"评价内容晒图"`
}
