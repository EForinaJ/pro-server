package dao_freelancer

type Edit struct {
	UId           string `json:"Uid"  dc:"Uid"`
	TitleId       int64  `json:"titleId"  dc:"头衔id"`
	CategoryId    int64  `json:"categoryId" dc:"分类id"`
	ContactMode   int    `json:"contactMode"  dc:"联系方式"`
	ContactNumber string `json:"contactNumber"  dc:"联系号码"`
}
