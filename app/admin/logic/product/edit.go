package product

import (
	"context"
	dto_product "server/app/admin/dto/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IProduct.
func (s *sProduct) Edit(ctx context.Context, req *dto_product.Edit) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var entity *do.SysProduct
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}

	entity.UpdateTime = gtime.Now()

	_, err = tx.Model(dao.SysProduct.Table()).
		Where(dao.SysProduct.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	// err = s.saveSpecAttrAndValue(tx, req.Id, req.Specification)
	// if err != nil {
	// 	return err
	// }
	// err = s.saveSkuAndRelated(tx, req.Id, req.Sku)
	// if err != nil {
	// 	return err
	// }

	return
}
