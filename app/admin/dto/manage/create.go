package dto_manage

type Create struct {
	Name     string  `p:"name"  v:"required#请设置管理员名称"  dc:"管理员名称"`
	Phone    string  `p:"phone"  v:"required|phone#请设置管理员手机号|手机号格式不正确"     dc:"手机号码"`
	Sex      int     `p:"sex"     dc:"性别"`
	Avatar   string  `p:"avatar"  dc:"头像地址"`
	Wechat   string  `p:"wechat"  dc:"微信二维码"`
	Password string  `p:"password"  v:"min-length:6#不能小于6位"  dc:"密码"`
	Status   int     `p:"status"    dc:"帐号状态（1停用,2正常）"`
	Roles    []int64 `p:"roles"   v:"required|array#请设置角色|数据格式为数组"    dc:"角色id"`
}
