package project

import (
	"context"
	dto_project "server/app/admin/dto/project"
)

// CheckDeductionMoney implements service.IProject.
func (s *sProject) CheckDeductionMoney(ctx context.Context, req *dto_project.Settlement) (res bool, err error) {
	// money, err := dao.SysProject.Ctx(ctx).Where(dao.SysProject.Columns().Id, req.Id).
	// 	Value(dao.SysProject.Columns().Money)
	// if err != nil {
	// 	return false, utils_error.Err(response.DB_READ_ERROR)
	// }
	// return decimal.NewFromFloat(money.Float64()).GreaterThan(decimal.NewFromFloat(req.Money)), nil

	return
}
