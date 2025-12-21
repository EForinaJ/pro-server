package product

import (
	dto_product "server/app/admin/dto/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sProduct) saveSpecAttrAndValue(tx gdb.TX, product int64, data []dto_product.Specification) (err error) {

	for _, v := range data {

		attrsId, err := tx.Model(dao.SysProductSpecAttrs.Table()).
			Where(dao.SysProductSpecAttrs.Columns().Name, v.Name).
			Where(dao.SysProductSpecAttrs.Columns().ProductId, product).
			Value(dao.SysProductSpecAttrs.Columns().Id)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		_, err = tx.Model(dao.SysProductSpecValues.Table()).
			Where(dao.SysProductSpecValues.Columns().SpecAttsId, attrsId).Delete()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		_, err = tx.Model(dao.SysProductSpecAttrs.Table()).
			Where(dao.SysProductSpecAttrs.Columns().Name, v.Name).
			Where(dao.SysProductSpecAttrs.Columns().ProductId, product).
			Delete()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		rs, err := tx.Model(dao.SysProductSpecAttrs.Table()).Data(g.Map{
			dao.SysProductSpecAttrs.Columns().Name:       v.Name,
			dao.SysProductSpecAttrs.Columns().ProductId:  product,
			dao.SysProductSpecAttrs.Columns().CreateTime: gtime.Now(),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		rid, err := rs.LastInsertId()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		for _, j := range v.Values {
			_, err = tx.Model(dao.SysProductSpecValues.Table()).Data(g.Map{
				dao.SysProductSpecValues.Columns().SpecAttsId: rid,
				dao.SysProductSpecValues.Columns().Value:      j,
				dao.SysProductSpecValues.Columns().CreateTime: gtime.Now(),
			}).Insert()
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}
		}
	}

	return
}
