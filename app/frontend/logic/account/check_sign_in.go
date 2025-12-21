package account

import (
	"context"
	"fmt"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// CheckSignIn implements service.IAccount.
func (s *sAccount) CheckSignIn(ctx context.Context) (err error) {
	now := time.Now()
	key := fmt.Sprintf("sign:%d:%s", ctx.Value("userId"), now.Format("200601")) // 例如: sign:123:202511
	offset := now.Day() - 1

	// 执行GETBIT命令
	bit, err := g.Redis().Do(ctx, "GETBIT", key, offset)
	if err != nil {
		return utils_error.Err(response.CACHE_READ_ERROR)
	}
	if bit.Int() == 1 {
		return utils_error.ErrMessage(response.FAILD, "已经签到了")
	}

	return
}
