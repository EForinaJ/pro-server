package dao_product

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id"     dc:"商品编号"`
	Name       string      `json:"name"   dc:"商品名称"`
	Pic        string      `json:"pic"  dc:"商品封面"`
	Game       string      `json:"game"   dc:"商品分类"`
	Price      float64     `json:"price"   dc:"商品售价"`
	Rate       int         `json:"rate" dc:"评分"`
	Category   string      `json:"category"   dc:"所属分类"`
	SalesCount int64       `json:"salesCount"   dc:"商品销量"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
	Status     int         `json:"status"   dc:"商品状态"`
}
