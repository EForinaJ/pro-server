package dict_data

import (
	"context"
	dao_dict_data "server/app/admin/dao/dict_data"
	dto_dict_data "server/app/admin/dto/dict_data"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDictData struct {
}

// GetAll implements service.IDictData.
func (s *sDictData) GetAll(ctx context.Context, code string) (res []*dao_dict_data.List, err error) {

	var list []*entity.SysDictData
	err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Code, code).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_dict_data.List, len(list))
	for i, v := range list {
		var entity *dao_dict_data.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}

// // Delete implements service.IDictData.
func (s *sDictData) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysDictData.Ctx(ctx).
		WhereIn(dao.SysDictData.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

// Edit implements service.IDictData.
func (s *sDictData) Edit(ctx context.Context, req *dto_dict_data.Edit) (err error) {
	var entity *do.SysDictData
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

// GetEdit implements service.IDictData.
func (s *sDictData) GetEdit(ctx context.Context, id int64) (res *dao_dict_data.Edit, err error) {
	err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// CheckUnique implements service.IDictData.
func (s *sDictData) CheckUnique(ctx context.Context, code, key string) (res bool, err error) {
	res, err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Key, key).
		Where(dao.SysDictData.Columns().Code, code).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// CheckUniqueById implements service.IDictData.
func (s *sDictData) CheckUniqueById(ctx context.Context, code, key string, id int64) (res bool, err error) {
	res, err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Key, key).
		WhereNotIn(dao.SysDictData.Columns().Id, id).
		Where(dao.SysDictData.Columns().Code, code).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// Create implements service.IDictData.
func (s *sDictData) Create(ctx context.Context, req *dto_dict_data.Create) (err error) {
	var entity *do.SysDictData
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysDictData.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}

func init() {
	service.RegisterDictData(&sDictData{})
}
