package dto_permission

type Create struct {
	PId         int64  `p:"pId" dc:"父id"`
	Permission  string `p:"permission" v:"required#请输入权限标识" dc:"权限标识"`
	Name        string `p:"name" v:"required#请输入权限名称" dc:"权限名称"`
	Description string `p:"description" dc:"权限描述"`
}
