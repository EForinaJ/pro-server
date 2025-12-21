// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure for table sys_comment.
type SysComment struct {
	Id          int64       `json:"id"          orm:"id"            description:""` //
	UserId      int64       `json:"userId"      orm:"user_id"       description:""` //
	OrderId     int64       `json:"orderId"     orm:"order_id"      description:""` //
	ProductId   int64       `json:"productId"   orm:"product_id"    description:""` //
	ParentId    int64       `json:"parentId"    orm:"parent_id"     description:""` //
	ReplyUserId int64       `json:"replyUserId" orm:"reply_user_id" description:""` //
	Content     string      `json:"content"     orm:"content"       description:""` //
	Rating      int         `json:"rating"      orm:"rating"        description:""` //
	Images      string      `json:"images"      orm:"images"        description:""` //
	IsTop       int         `json:"isTop"       orm:"is_top"        description:""` //
	Status      int         `json:"status"      orm:"status"        description:""` //
	Ip          string      `json:"ip"          orm:"ip"            description:""` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:""` //
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   description:""` //
}
