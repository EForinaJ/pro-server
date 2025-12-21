package product

import (
	"context"
	dto_product "server/app/admin/dto/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

// CheckSkuEmpty implements service.IProduct.
func (s *sProduct) CheckSkuEmpty(ctx context.Context, req []dto_product.Sku) (err error) {
	if len(req) == 0 {
		return utils_error.ErrMessage(response.FAILD, "请设置规格")
	}
	for _, v := range req {
		if v.Price == 0 {
			return utils_error.ErrMessage(response.FAILD, "请设置规格价格")
		}
	}
	return nil
}
