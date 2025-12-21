package product

import (
	dto_product "server/app/admin/dto/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sProduct) saveSkuAndRelated(tx gdb.TX, product int64, data []dto_product.Sku) (err error) {
	_, err = tx.Model(dao.SysProductSku.Table()).
		Where(dao.SysProductSku.Columns().ProductId, product).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	for _, v := range data {
		var entity *do.SysProductSku
		err = gconv.Scan(v, &entity)
		if err != nil {
			return utils_error.Err(response.FAILD)
		}
		entity.ProductId = product
		entity.CreatedTime = gtime.Now()
		entity.UpdatedTime = gtime.Now()

		rs, err := tx.Model(dao.SysProductSku.Table()).
			Data(entity).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		rid, err := rs.LastInsertId()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		_, err = tx.Model(dao.SysProductSkuSpecRelations.Table()).
			Where(dao.SysProductSkuSpecRelations.Columns().SkuId, rid).
			Delete()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		for _, attr := range v.Attrs {
			attrId, err := tx.Model(dao.SysProductSpecAttrs.Table()).
				Where(dao.SysProductSpecAttrs.Columns().ProductId, product).
				Where(dao.SysProductSpecAttrs.Columns().Name, attr.Attr).
				Value(dao.SysProductSpecAttrs.Columns().Id)
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}

			valueId, err := tx.Model(dao.SysProductSpecValues.Table()).
				Where(dao.SysProductSpecValues.Columns().SpecAttsId, attrId).
				Where(dao.SysProductSpecValues.Columns().Value, attr.Value).
				Value(dao.SysProductSpecValues.Columns().Id)
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}

			_, err = tx.Model(dao.SysProductSkuSpecRelations.Table()).
				Data(g.Map{
					dao.SysProductSkuSpecRelations.Columns().SpecAttrId:  attrId,
					dao.SysProductSkuSpecRelations.Columns().SpecValueId: valueId,
					dao.SysProductSkuSpecRelations.Columns().SkuId:       rid,
					dao.SysProductSkuSpecRelations.Columns().CreateTime:  gtime.Now(),
				}).Insert()
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}
		}

	}
	return
}
