package product

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IProduct.
func (s *sProduct) Delete(ctx context.Context, ids []int64) (err error) {
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

	// skus, err := tx.Model(dao.SysProductSku.Table()).
	// 	Fields(dao.SysProductSku.Columns().Id).
	// 	WhereIn(dao.SysProductSku.Columns().ProductId, ids).Array()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSkuSpecRelations.Table()).
	// 	WhereIn(dao.SysProductSkuSpecRelations.Columns().SkuId, skus).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// attrs, err := tx.Model(dao.SysProductSpecAttrs.Table()).
	// 	Fields(dao.SysProductSpecAttrs.Columns().Id).
	// 	WhereIn(dao.SysProductSpecAttrs.Columns().ProductId, ids).Array()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSpecValues.Table()).
	// 	WhereIn(dao.SysProductSpecValues.Columns().SpecAttsId, attrs).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSpecAttrs.Table()).
	// 	WhereIn(dao.SysProductSpecAttrs.Columns().ProductId, ids).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSku.Table()).
	// 	WhereIn(dao.SysProductSku.Columns().ProductId, ids).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	_, err = tx.Model(dao.SysProduct.Table()).
		WhereIn(dao.SysProduct.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
