package dto_distribute

type Cancel struct {
	Id     int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"分类id"`
	Reason string `p:"reason" v:"required#请输入取消原因" dc:"取消原因"`
}
