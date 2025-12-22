package dto_order

type ChangeWitkey struct {
	Id     int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	OldId  int64  `p:"oldId" v:"required#请设置变更威客ID" dc:"变更威客ID"`
	NewId  int64  `p:"newId" v:"required#请设置新威客ID" dc:"新威客ID"`
	Reason string `p:"reason" v:"required#请输入变更原因" dc:"变更原因"`
}
