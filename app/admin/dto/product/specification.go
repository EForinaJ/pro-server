package dto_product

type Specification struct {
	Name   string   `p:"name" dc:"规格名称"`
	Values []string `p:"values" dc:"规格值"`
}
