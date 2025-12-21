package casbin

import (
	"context"
	"server/internal/dao"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gconv"
)

var Enforcer *casbin.Enforcer

// InitEnforcer 初始化
func InitEnforcer(ctx context.Context) {
	var (
		link   = g.Cfg().MustGet(ctx, "database.default.link")
		a, err = NewAdapter(link.String())
	)

	path := "./config/casbin.conf"

	// 优先从本地加载casbin.conf，如果不存在就从资源文件中找
	modelContent := gfile.GetContents(path)
	if len(modelContent) == 0 {
		if !gres.IsEmpty() && gres.Contains(path) {
			modelContent = string(gres.GetContent(path))
		}
	}

	if len(modelContent) == 0 {
		g.Log().Panicf(ctx, "casbin model file does not exist：%v", path)
	}

	m, err := model.NewModelFromString(modelContent)
	if err != nil {
		g.Log().Panicf(ctx, "casbin NewModelFromString err：%v", err)
	}

	Enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		g.Log().Panicf(ctx, "casbin NewEnforcer err：%v", err)
	}

	loadPermissions(ctx)
}

func loadPermissions(ctx context.Context) {
	// Refresh(ctx)
	// 获取角色权限
	rs, err := dao.SysRolePermission.Ctx(ctx).
		As("a").
		LeftJoin("sys_role", "b", "b.id=a.role_id").
		LeftJoin("sys_permission", "c", "c.id=a.permission_id").
		Fields("c.permission", "b.code").
		All()
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}

	rules := make([][]string, len(rs.List()))
	for i, v := range rs.List() {
		entites := make([]string, 2)
		entites[0] = gconv.String(v["code"])
		entites[1] = gconv.String(v["permission"])
		rules[i] = entites
	}
	if _, err = Enforcer.AddPolicies(rules); err != nil {
		g.Log().Fatalf(ctx, "loadPermissions AddPolicies err:%v", err)
		return
	}

	//  添加和角色管理表内容到casbin
	// 获取角色权限
	mrs, err := dao.SysManageRole.Ctx(ctx).
		As("a").
		LeftJoin("sys_role", "b", "b.id=a.role_id").
		LeftJoin("sys_manage", "c", "c.id=a.manage_id").
		Fields("c.id", "b.code").
		All()
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}
	roles := make([][]string, len(mrs.List()))
	for i, v := range mrs.List() {
		entites := []string{
			gconv.String(v["id"]),
			gconv.String(v["code"]),
		}
		roles[i] = entites
	}
	if _, err = Enforcer.AddGroupingPolicies(roles); err != nil {
		g.Log().Fatalf(ctx, "loadPermissions AddPolicies err:%v", err)
		return
	}

}
func Clear(ctx context.Context) (err error) {
	policy, err := Enforcer.GetPolicy()
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}
	_, err = Enforcer.RemovePolicies(policy)
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}

	// 检查是否清理干净
	if len(policy) > 0 {
		return Clear(ctx)
	}
	return
}

func Refresh(ctx context.Context) (err error) {
	if err = Clear(ctx); err != nil {
		return err
	}
	loadPermissions(ctx)
	return
}
