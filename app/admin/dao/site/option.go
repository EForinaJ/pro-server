package dao_site

type Options struct {
	Id    int64       `json:"id" dc:"ID"`
	Value interface{} `json:"value" dc:"值"`
	Name  string      `json:"name" dc:"昵称"`
}
