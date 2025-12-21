package product

import (
	"context"
	dao_product "server/app/admin/dao/product"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetEdit implements service.IProduct.
func (s *sProduct) GetEdit(ctx context.Context, id int64) (res *dao_product.Edit, err error) {
	err = dao.SysProduct.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 获取所有规格属性
	// attrs, err := dao.SysProductSpecAttrs.Ctx(ctx).
	// 	Where(dao.SysProductSpecAttrs.Columns().ProductId, id).All()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// specList := make([]dao_product.Specification, attrs.Len())
	// for iAttr, vAttr := range attrs {
	// 	var spec dao_product.Specification
	// 	spec.Name = vAttr.GMap().Get(dao.SysProductSpecAttrs.Columns().Name).(string)

	// 	values, err := dao.SysProductSpecValues.Ctx(ctx).
	// 		Fields(dao.SysProductSpecValues.Columns().Value).
	// 		Where(dao.SysProductSpecValues.Columns().SpecAttsId, vAttr.GMap().Get(dao.SysProductSpecAttrs.Columns().Id).(int64)).
	// 		Array()
	// 	if err != nil {
	// 		return nil, utils_error.Err(response.DB_READ_ERROR)
	// 	}
	// 	spec.Values = gconv.Strings(values)

	// 	specList[iAttr] = spec
	// }
	// res.Specification = specList

	// //  获取规格
	// skus, err := dao.SysProductSku.Ctx(ctx).
	// 	Where(dao.SysProductSku.Columns().ProductId, id).All()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// skuList := make([]dao_product.Sku, skus.Len())
	// for sSku, vSku := range skus {
	// 	var sku dao_product.Sku
	// 	sku.Code = gconv.String(vSku.GMap().Get(dao.SysProductSku.Columns().Code))
	// 	sku.Price = gconv.Float64(vSku.GMap().Get(dao.SysProductSku.Columns().Price))
	// 	sku.OriginalPrice = gconv.Float64(vSku.GMap().Get(dao.SysProductSku.Columns().OriginalPrice))
	// 	sku.Stock = gconv.Int(vSku.GMap().Get(dao.SysProductSku.Columns().Stock))

	// 	attrs := gstr.Split(sku.Code, "|")
	// 	sku.AttrNames = make([]string, len(attrs))
	// 	sku.Values = make([]string, len(attrs))
	// 	sku.Attrs = make([]dao_product.Attrs, len(attrs))
	// 	for iAttr, vAttr := range attrs {
	// 		attrValue := gstr.Split(vAttr, ":")
	// 		sku.Values[iAttr] = attrValue[1]
	// 		sku.AttrNames[iAttr] = attrValue[0]
	// 		sku.Attrs[iAttr] = dao_product.Attrs{
	// 			Attr:  attrValue[0],
	// 			Value: attrValue[1],
	// 		}
	// 	}

	// 	skuList[sSku] = sku
	// }
	// res.Sku = skuList
	return
}
