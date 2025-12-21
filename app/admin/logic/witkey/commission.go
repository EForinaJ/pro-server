package witkey

import (
	"context"
	dto_witkey "server/app/admin/dto/witkey"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// ChangeCommission implements service.IWitkey.
func (s *sWitkey) ChangeCommission(ctx context.Context, req *dto_witkey.Commission) (err error) {
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
	commission, err := tx.Model(dao.SysWitkey.Table()).
		WherePri(req.Id).Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	amount := decimal.NewFromFloat(commission.Float64())
	entity := g.Map{
		dao.SysCommission.Columns().After:      commission,
		dao.SysCommission.Columns().Money:      req.Money,
		dao.SysCommission.Columns().Mode:       req.Mode,
		dao.SysCommission.Columns().WitkeyId:   req.Id,
		dao.SysCommission.Columns().CreateTime: gtime.Now(),
		dao.SysCommission.Columns().Type:       consts.UserChangeCommissionTypeSystem,
	}

	switch req.Mode {
	case consts.FundLogModeAdd:
		newCommission := amount.Add(decimal.NewFromFloat(req.Money))
		_, err = tx.Model(dao.SysWitkey.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysWitkey.Columns().Commission: newCommission,
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity[dao.SysCommission.Columns().Before] = newCommission
		entity[dao.SysCommission.Columns().Remark] = "系统增加余额"
		entity[dao.SysCommission.Columns().Related] = "系统增加余额"
	case consts.FundLogModeSub:
		newCommission := amount.Sub(decimal.NewFromFloat(req.Money))
		_, err = tx.Model(dao.SysWitkey.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysWitkey.Columns().Commission: newCommission,
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity[dao.SysCommission.Columns().Related] = "系统减少余额"
		entity[dao.SysCommission.Columns().Remark] = "系统减少余额"
		entity[dao.SysCommission.Columns().Before] = newCommission
	}

	if req.Remark != "" {
		entity[dao.SysCommission.Columns().Remark] = req.Remark
	}

	_, err = tx.Model(dao.SysCommission.Table()).
		Data(entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
