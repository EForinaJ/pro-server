package account

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_account "server/app/frontend/dao/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetWorkspace implements service.IAccount.
func (s *sAccount) GetWorkspace(ctx context.Context) (res *dao_account.Workspace, err error) {
	var entity dao_account.Workspace

	witkey, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, ctx.Value("userId")).
		Value(dao.SysWitkey.Columns().CreateTime)
	if err != nil {
		return nil, utils_error.Err(response.FAILD)
	}
	entity.SettleIn = int(gtime.Now().Sub(witkey.GTime()).Hours() / 24)

	commission, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().WitkeyId, ctx.Value("userId")).
		// Where(dao.SysProject.Columns().Status, consts.ProjectStatusCompleted).
		// Where(dao.SysProject.Columns().IsSettlement, consts.Yes).
		Sum("commission")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	entity.TotalRevenue = commission

	projectCount, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().WitkeyId, ctx.Value("userId")).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	entity.ProjectCount = projectCount

	afterSalesCount, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().WitkeyId, ctx.Value("userId")).
		Where(dao.SysProject.Columns().Status, consts.ProjectStatusCancel).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	entity.AfterSalesCount = afterSalesCount

	// announcement, err := dao.SysAnnouncement.Ctx(ctx).
	// 	Page(1, 5).
	// 	OrderDesc(dao.SysProject.Columns().CreateTime).
	// 	All()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// announcementList := make([]*dao_account.Announcement, announcement.Len())
	// for i, v := range announcement {
	// 	var announcement *dao_account.Announcement
	// 	err = gconv.Scan(v.GMap(), &announcement)
	// 	if err != nil {
	// 		return nil, utils_error.Err(response.DB_READ_ERROR)
	// 	}
	// 	announcementList[i] = announcement
	// }
	// entity.Announcement = announcementList

	project, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().WitkeyId, ctx.Value("userId")).
		OrderDesc(dao.SysProject.Columns().CreateTime).
		Fields(dao.SysProject.Columns().CreateTime,
			dao.SysProject.Columns().Id,
			// dao.SysProject.Columns().Code,
			dao.SysProject.Columns().Status).
		Page(1, 5).
		All()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	projectList := make([]*dao_account.Project, project.Len())
	for i, v := range project {
		var project *dao_account.Project
		err = gconv.Scan(v.GMap(), &project)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		projectList[i] = project
	}
	entity.Project = projectList

	contact, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.ContactSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// contact, err := service.System().GetOne(ctx, consts.ContactSetting)
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }

	contactJson, err := gjson.DecodeToJson(contact)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	var contactDetail *dao_account.Contact
	err = contactJson.Scan(&contactDetail)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	entity.Contact = contactDetail

	//  获取客服

	roleId, err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Code, contactJson.Get("role")).
		Value(dao.SysRole.Columns().Id)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	manageIds, err := dao.SysManageRole.Ctx(ctx).
		Where(dao.SysManageRole.Columns().RoleId, roleId).
		Fields(dao.SysManageRole.Columns().ManageId).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	manages, err := dao.SysManage.Ctx(ctx).
		WhereIn(dao.SysManage.Columns().Id, manageIds).
		Fields(
			dao.SysManage.Columns().Id,
			dao.SysManage.Columns().Avatar,
			dao.SysManage.Columns().Name,
			// dao.SysManage.Columns().Description,
			// dao.SysManage.Columns().WxNumber,
		).
		All()
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	managesList := make([]*dao_account.Manage, manages.Len())
	for i, v := range manages {
		var manage *dao_account.Manage
		err = gconv.Scan(v.GMap(), &manage)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		managesList[i] = manage
	}
	entity.Manage = managesList

	return &entity, nil
}
