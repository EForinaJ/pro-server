package dao_project

type Order struct {
	Id           int64    `json:"id" dc:"ID"`
	Code         string   `json:"code" dc:"订单号"`
	Product      *Product `json:"product" dc:"下单商品"`
	Requirements string   `json:"requirements" dc:"订单要求"`
	Status       int      `json:"status" dc:"订单状态"`
}
