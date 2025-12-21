package dao_dict_data

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id"     dc:"字典编号"`
	Name       string      `json:"name"   dc:"字典数据名称"`
	Key        string      `json:"key"   dc:"字典数据键"`
	Value      string      `json:"value" dc:"字典数据值"`
	Code       string      `json:"code"   dc:"字典码"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
