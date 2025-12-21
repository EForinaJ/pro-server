package dto_role

type Create struct {
	Name        string `p:"name" v:"required#请输入角色名称" dc:"角色名称"`
	Code        string `p:"code" v:"required#请输入角色编码" dc:"角色编码"`
	Description string `p:"description" dc:"角色描述"`
	Status      int    `p:"status" v:"required|in:1,2#请选择角色状态|角色状态只允许1或2" dc:"角色状态"`
}
