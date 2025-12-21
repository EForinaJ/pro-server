package dto_role

type Edit struct {
	Id          int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"角色id"`
	Name        string `p:"name" v:"required#请输入角色名称" dc:"角色名称"`
	Code        string `p:"code" v:"required#请输入角色编码" dc:"角色编码"`
	Description string `p:"description" dc:"角色描述"`
	Status      int    `p:"status" v:"required|in:1,2#请选择角色状态|角色状态只允许1或2" dc:"角色状态"`
}
