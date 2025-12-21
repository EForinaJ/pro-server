package product

import (
	"context"
	dao_product "server/app/admin/dao/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IPunish.
func (s *sProduct) GetAllByName(ctx context.Context, name string) (res []*dao_product.List, err error) {
	m := dao.SysProduct.Ctx(ctx).
		OrderDesc(dao.SysProduct.Columns().CreateTime)
	m = m.WhereLike(dao.SysProduct.Columns().Name, "%"+name+"%")

	var list []*entity.SysProduct
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_product.List, len(list))
	for i, v := range list {
		var entity *dao_product.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
