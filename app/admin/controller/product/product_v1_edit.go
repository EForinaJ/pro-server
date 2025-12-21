package product

import (
	"context"

	v1 "server/app/admin/api/product/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	// err = service.Product().CheckSpecAttrEmpty(ctx, req.Specification)
	// if err != nil {
	// 	return nil, err
	// }

	// err = service.Product().CheckSkuEmpty(ctx, req.Sku)
	// if err != nil {
	// 	return nil, err
	// }
	err = service.Product().Edit(ctx, req.Edit)
	return
}
