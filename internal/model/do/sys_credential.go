// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCredential is the golang structure of table sys_credential for DAO operations like Where/Data.
type SysCredential struct {
	g.Meta                 `orm:"table:sys_credential, do:true"`
	Id                     interface{} //
	UserId                 interface{} //
	CredentialType         interface{} //
	CredentialName         interface{} //
	CredentialNumber       interface{} //
	CredentialFrontImage   interface{} //
	CredentialAgainstImage interface{} //
	Status                 interface{} //
	CreateTime             *gtime.Time //
	UpdateTime             *gtime.Time //
}
