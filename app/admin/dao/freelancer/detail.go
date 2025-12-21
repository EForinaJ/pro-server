package dao_freelancer

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id            int64       `json:"id" dc:"ID"`
	Category      string      `json:"category" dc:"类目"`
	Title         string      `json:"title" dc:"头衔"`
	Examiner      string      `json:"examiner" dc:"考官"`
	Name          string      `json:"name" dc:"申请人"`
	ContactMode   int         `json:"contactMode" dc:"联系方式"`
	ContactNumber string      `json:"contactNumber" dc:"联系号码"`
	Images        []string    `json:"images" dc:"段位截图"`
	CreateTime    *gtime.Time `json:"createTime" dc:"申请时间"`
	Status        int         `json:"status" dc:"申请状态"`
	Remark        string      `json:"remark" dc:"备注"`
}
