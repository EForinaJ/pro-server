package dao_product

type Specification struct {
	Name   string   `json:"name" dc:"规格名称"`
	Values []string `json:"values" dc:"规格值"`
}
