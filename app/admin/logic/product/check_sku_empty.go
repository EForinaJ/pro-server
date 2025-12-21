package product

import (
	"context"
	dto_product "server/app/admin/dto/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

// CheckSpecAttrEmpty implements service.IProduct.
func (s *sProduct) CheckSpecAttrEmpty(ctx context.Context, req []dto_product.Specification) (err error) {
	for _, v := range req {
		if v.Name == "" {
			return utils_error.ErrMessage(response.FAILD, "请输入规格属性名称")
		}
		for _, j := range v.Values {
			if j == "" {
				return utils_error.ErrMessage(response.FAILD, "请输入对应值")
			}
		}
	}
	return nil
}
