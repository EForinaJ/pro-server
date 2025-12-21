package dao_dict_type

type Edit struct {
	Id   int64  `json:"id"     dc:"字典编号"`
	Name string `json:"name"   dc:"字典名称"`
	Code string `json:"code"   dc:"字典码"`
}
