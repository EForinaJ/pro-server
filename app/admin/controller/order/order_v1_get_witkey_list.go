package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetWitkeyList(ctx context.Context, req *v1.GetWitkeyListReq) (res *v1.GetWitkeyListRes, err error) {
	total, list, err := service.Order().GetWitkeyList(ctx, req.WitkeyQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetWitkeyListRes{
		Total: total,
		List:  list,
	}

	return
}
