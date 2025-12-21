package utils_snow

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

func GetCode(ctx context.Context, key string) string {
	return key + gtime.Now().Format("YnjHis") + gconv.String(grand.Intn(999999)) + gconv.String(grand.Intn(99))
}
