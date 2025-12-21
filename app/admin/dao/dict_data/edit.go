package dao_dict_data

type Edit struct {
	Id    int64  `json:"id"     dc:"字典编号"`
	Name  string `json:"name"   dc:"字典数据名称"`
	Key   string `json:"key"   dc:"字典数据名键"`
	Value string `json:"value" dc:"字典数据值"`
	Code  string `json:"code"   dc:"字典码"`
}
