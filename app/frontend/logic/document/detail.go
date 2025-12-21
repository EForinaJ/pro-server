package document

import (
	"context"
	dao_document "server/app/frontend/dao/document"
)

// GetDetail implements service.IGoods.
func (s *sDocument) GetDetail(ctx context.Context, id int64) (res *dao_document.Detail, err error) {
	// detail, err := dao.SysDocument.Ctx(ctx).WherePri(id).One()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }

	// err = gconv.Scan(detail, &res)
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	return
}
