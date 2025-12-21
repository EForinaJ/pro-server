// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"server/app/frontend/api/auth/v1"
)

type IAuthV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	ProgramLogin(ctx context.Context, req *v1.ProgramLoginReq) (res *v1.ProgramLoginRes, err error)
}
