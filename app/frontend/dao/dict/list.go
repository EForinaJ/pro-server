package dao_dict

type List struct {
	Id    int64  `json:"id"     dc:"字典编号"`
	Name  string `json:"name" dc:"字典名称"`
	Value string `json:"value" dc:"字典值"`
}
