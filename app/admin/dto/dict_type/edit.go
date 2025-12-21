package dto_dict_type

type Edit struct {
	Id   int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Name string `p:"name" v:"required#请输入名称" dc:"名称"`
	Code string `p:"code" v:"required#请输入字典码" dc:"字典码"`
}
