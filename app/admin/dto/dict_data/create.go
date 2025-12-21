package dto_dict_data

type Create struct {
	Name  string `p:"name" v:"required#请输入字典数据名称" dc:"字典数据名称"`
	Key   string `p:"key" v:"required#请输入字典数据键" dc:"字典数据键"`
	Value string `p:"value" v:"required#请输入字典数据值" dc:"字典数据值"`
	Code  string `p:"code" v:"required#请输入字典码" dc:"字典码"`
}
