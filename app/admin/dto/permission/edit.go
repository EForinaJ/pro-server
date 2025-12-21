package dto_permission

type Edit struct {
	Id          int64  `p:"id" v:"required|integer|min:1#请输入菜单id|id类型必须是整型|id最小为1" dc:"菜单id"`
	PId         int64  `p:"pId" dc:"父id"`
	Permission  string `p:"permission" v:"required#请输入权限标识" dc:"权限标识"`
	Name        string `p:"name" v:"required#请输入权限名称" dc:"权限名称"`
	Description string `p:"description" dc:"权限描述"`
}
