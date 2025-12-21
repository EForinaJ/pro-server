package service

import (
	"context"
	dao_bill "server/app/frontend/dao/bill"
	dto_bill "server/app/frontend/dto/bill"
)

// 定义显示接口
type IBill interface {
	GetList(ctx context.Context, req *dto_bill.Query) (total int, res []*dao_bill.List, err error)
}

// 定义接口变量
var localBill IBill

// 定义一个获取接口的方法
func Bill() IBill {
	if localBill == nil {
		panic("implement not found for interface IBill, forgot register?")
	}
	return localBill
}

// 定义一个接口实现的注册方法
func RegisterBill(i IBill) {
	localBill = i
}
