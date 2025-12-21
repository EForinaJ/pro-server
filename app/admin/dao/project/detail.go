package dao_project

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id           int64       `json:"id" dc:"ID"`
	Code         string      `json:"code" dc:"项目号"`
	Witkey       *Witkey     `json:"witkey" dc:"接单者"`
	Order        *Order      `json:"order" dc:"所属订单"`
	Money        float64     `json:"money" dc:"项目金额"`
	ServiceFee   float64     `json:"serviceFee" dc:"平台抽成"`
	Commission   float64     `json:"commission" dc:"结算佣金"`
	Images       []string    `json:"images" dc:"项目截图"`
	StartTime    *gtime.Time `json:"startTime" dc:"服务开始时间"`
	FinishTime   *gtime.Time `json:"finishTime" dc:"服务完成时间"`
	CreateTime   *gtime.Time `json:"createTime" dc:"项目时间"`
	Status       int         `json:"status" dc:"项目状态"`
	IsSettlement int         `json:"isSettlement"   dc:"结算状态"`
	Remark       string      `json:"remark" dc:"管理备注"`
}
