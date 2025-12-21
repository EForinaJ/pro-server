package dto_dict_data

type Edit struct {
	Id    int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Key   string `p:"key" v:"required#请输入字典数据键" dc:"字典数据键"`
	Value string `p:"value" v:"required#请输入字典数据值" dc:"字典数据值"`
	Code  string `p:"code" v:"required#请输入字典码" dc:"字典码"`
	Name  string `p:"name" v:"required#请输入字典数据名称" dc:"字典数据名称"`
}
