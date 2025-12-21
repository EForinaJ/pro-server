package v1

import (
	dao_site "server/app/admin/dao/site"

	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------- 基础配置
type GetSiteReq struct {
	g.Meta `path:"/site" method:"get" tags:"站点" summary:"获取系统站点信息"`
}
type GetSiteRes struct {
	*dao_site.Detail
}

type GetUserOptionsReq struct {
	g.Meta `path:"/site/user/options" method:"get" tags:"站点" summary:"所属用户选项"`
	Name   string `p:"name" v:"required#请输入用户昵称" dc:"用户昵称"`
}
type GetUserOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"所属用户选项列表"`
}

type GetGameOptionsReq struct {
	g.Meta `path:"/site/game/options" method:"get" tags:"站点" summary:"游戏选项"`
}
type GetGameOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"游戏选项列表"`
}

type GetCategoryeOptionsReq struct {
	g.Meta `path:"/site/category/options" method:"get" tags:"站点" summary:"商品分类选项"`
	GameId int64 `p:"gameId" v:"required#请输入游戏id" dc:"游戏id"`
}
type GetCategoryeOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"商品分类选项列表"`
}

type GetTitleOptionsReq struct {
	g.Meta `path:"/site/title/options" method:"get" tags:"站点" summary:"所属头衔选项"`
	GameId int64 `p:"gameId" v:"required#请输入游戏id" dc:"游戏id"`
}
type GetTitleOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"所属头衔选项列表"`
}
