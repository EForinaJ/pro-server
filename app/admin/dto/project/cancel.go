package dto_project

type Cancel struct {
	Id      int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Content string `p:"conetnt" v:"required#请输入取消原因" dc:"取消原因"`
}
