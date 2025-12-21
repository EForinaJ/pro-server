package dict_type

import (
	"context"
	dao_dict_type "server/app/admin/dao/dict_type"
	dto_dict_type "server/app/admin/dto/dict_type"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDictType struct {
}

// Delete implements service.IDictType.
func (s *sDictType) Delete(ctx context.Context, ids []int64) (err error) {
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

	_, err = tx.Model(dao.SysDictType.Table()).
		WhereIn(dao.SysDictType.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	codes, err := tx.Model(dao.SysDictType.Table()).WhereIn(dao.SysDictType.Columns().Id, ids).
		Fields(dao.SysDictType.Columns().Code).Array()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	_, err = tx.Model(dao.SysDictData.Table()).
		WhereIn(dao.SysDictData.Columns().Code, codes).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}

// Edit implements service.IDictType.
func (s *sDictType) Edit(ctx context.Context, req *dto_dict_type.Edit) (err error) {
	var entity *do.SysDictType
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysDictType.Ctx(ctx).
		Where(dao.SysDictType.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

// GetEdit implements service.IDictType.
func (s *sDictType) GetEdit(ctx context.Context, id int64) (res *dao_dict_type.Edit, err error) {
	err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// CheckUnique implements service.IDictType.
func (s *sDictType) CheckUnique(ctx context.Context, code string) (res bool, err error) {
	res, err = dao.SysDictType.Ctx(ctx).
		Where(dao.SysDictType.Columns().Code, code).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// CheckUniqueById implements service.IDictType.
func (s *sDictType) CheckUniqueById(ctx context.Context, code string, id int64) (res bool, err error) {
	res, err = dao.SysDictType.Ctx(ctx).
		WhereNotIn(dao.SysDictType.Columns().Id, id).
		Where(dao.SysDictType.Columns().Code, code).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// Create implements service.IDictType.
func (s *sDictType) Create(ctx context.Context, req *dto_dict_type.Create) (err error) {
	var entity *do.SysDictType
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysDictType.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}

// GetList implements service.IDictType.
func (s *sDictType) GetList(ctx context.Context, req *dto_dict_type.Query) (total int, res []*dao_dict_type.List, err error) {
	m := dao.SysDictType.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysDictType.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysDictType.Columns().Name, req.Name)
	}

	if req.Code != "" {
		m = m.Where(dao.SysDictType.Columns().Code, req.Code)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	err = m.Scan(&res)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	return
}

func init() {
	service.RegisterDictType(&sDictType{})
}
