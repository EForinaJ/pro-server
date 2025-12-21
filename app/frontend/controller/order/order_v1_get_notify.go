package order

import (
	"context"

	v1 "server/app/frontend/api/order/v1"
	"server/app/frontend/service"

	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) GetNotify(ctx context.Context, req *v1.GetNotifyReq) (res *v1.GetNotifyRes, err error) {
	err = service.Order().Notify(ctx)
	if err != nil {
		glog.Info(ctx, "微信支付结果异常错误")
		return
	}
	return
}
