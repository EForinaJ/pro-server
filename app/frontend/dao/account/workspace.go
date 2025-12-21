package dao_account

type Workspace struct {
	SettleIn        int                `json:"settleIn" dc:"入驻天数"`
	TotalRevenue    float64            `json:"totalRevenue" dc:"总收益"`
	ProjectCount    int                `json:"projectCount" dc:"项目接单总数"`
	AfterSalesCount int                `json:"afterSalesCount" dc:"售后取消总数"`
	Contact         *Contact           `json:"contact" dc:"平台联系"`
	CustomerService []*CustomerService `json:"customerService" dc:"客服列表"`
	Project         []*Project         `json:"project" dc:"项目"`
	Announcement    []*Announcement    `json:"announcement" dc:"公告列表"`
	Manage          []*Manage          `json:"manage" dc:"客服管理"`
}

type Contact struct {
	Platform string `json:"platform" dc:"联系平台"`
	Number   string `json:"number" dc:"平台号吗"`
	Icon     string `json:"icon" dc:"平台图标"`
}

type CustomerService struct {
	Avatar      string `json:"avatar" dc:"头像"`
	Name        string `json:"name" dc:"名称"`
	Description string `json:"description" dc:"描述"`
	WxNumber    string `json:"WxNumber" dc:"微信号"`
}

type Announcement struct {
	Id         int64  `json:"id" dc:"ID"`
	Type       int    `json:"type" dc:"公告类型"`
	Title      string `json:"title" dc:"公告标题"`
	CreateTime string `json:"createTime" dc:"发布时间"`
}

type Project struct {
	Id         int64  `json:"id" dc:"ID"`
	Code       string `json:"code" dc:"项目编号"`
	Status     int    `json:"status" dc:"项目状态"`
	CreateTime string `json:"createTime" dc:"接单时间"`
}

type Manage struct {
	Id          int64  `json:"id" dc:"ID"`
	Name        string `json:"name" dc:"名称"`
	Avatar      string `json:"avatar" dc:"头像"`
	Description string `json:"description" dc:"描述"`
	WxNumber    string `json:"wxNumber" dc:"微信号"`
}
