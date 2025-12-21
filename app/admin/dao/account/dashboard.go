package dao_account

type Dashboard struct {
	HotGoods                  interface{} `json:"hotGoods" dc:"热门商品"`
	CountCard                 interface{} `json:"countCard" dc:"数量卡片"`
	OrderStatistics           interface{} `json:"orderStatistics" dc:"订单统计"`
	UserConsumptionStatistics interface{} `json:"userConsumptionStatistics" dc:"用户消费统计"`
	LineOrderStatistics       interface{} `json:"lineOrderStatistics" dc:"用户消费统计"`
}
