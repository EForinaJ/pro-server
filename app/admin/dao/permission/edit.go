package dao_permission

type Edit struct {
	Id          int64  `json:"id" dc:"权限ID"`
	PId         int64  `json:"pId" dc:"父Id"`
	Permission  string `json:"permission" dc:"权限标识"`
	Name        string `json:"name" dc:"权限名称"`
	Description string `json:"description" dc:"权限描述"`
}
