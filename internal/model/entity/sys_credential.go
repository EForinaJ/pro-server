// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCredential is the golang structure for table sys_credential.
type SysCredential struct {
	Id                     int64       `json:"id"                     orm:"id"                       description:""` //
	UserId                 int64       `json:"userId"                 orm:"user_id"                  description:""` //
	CredentialType         int         `json:"credentialType"         orm:"credential_type"          description:""` //
	CredentialName         string      `json:"credentialName"         orm:"credential_name"          description:""` //
	CredentialNumber       string      `json:"credentialNumber"       orm:"credential_number"        description:""` //
	CredentialFrontImage   string      `json:"credentialFrontImage"   orm:"credential_front_image"   description:""` //
	CredentialAgainstImage string      `json:"credentialAgainstImage" orm:"credential_against_image" description:""` //
	Status                 int         `json:"status"                 orm:"status"                   description:""` //
	CreateTime             *gtime.Time `json:"createTime"             orm:"create_time"              description:""` //
	UpdateTime             *gtime.Time `json:"updateTime"             orm:"update_time"              description:""` //
}
