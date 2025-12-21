package v1

import (
	dao_site "server/app/frontend/dao/site"

	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------- 基础配置
type GetSiteReq struct {
	g.Meta `path:"/site" method:"get" tags:"站点" summary:"获取系统站点信息"`
}
type GetSiteRes struct {
	*dao_site.Detail
}

// ---------------------- 字典信息
type GetSiteDictDataReq struct {
	g.Meta `path:"/site/dict/data" method:"get" tags:"站点" summary:"获取系统站点字典信息"`
	Code   string `p:"code" v:"required#请设置字典数据码" dc:"字典数据码"`
}
type GetSiteDictDataRes struct {
	List []*dao_site.DictData `json:"list" dc:"字典数据列表"`
}
