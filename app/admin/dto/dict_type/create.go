package dto_dict_type

type Create struct {
	Name string `p:"name" v:"required#请输入名称" dc:"名称"`
	Code string `p:"code" v:"required#请输入字典码" dc:"字典码"`
}
