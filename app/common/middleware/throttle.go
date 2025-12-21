package middleware

// slidingWindowLimit 滑动窗口限流函数
// key: 限流键
// windowSize: 窗口大小，如 time.Minute
// maxRequests: 窗口内允许的最大请求数
// func slidingWindowLimit(ctx context.Context, key string, windowSize time.Duration, maxRequests int64) (bool, error) {
// 	now := time.Now().UnixMilli()
// 	windowStart := now - windowSize.Milliseconds()

// 	// 1. 移除窗口之前的旧数据
// 	_, err := g.Redis().ZRemRangeByScore(ctx, key, "0", fmt.Sprint(windowStart))
// 	if err != nil {
// 		return false, err
// 	}

// 	// 2. 获取当前窗口内的请求数
// 	count, err := g.Redis().ZCard(ctx, key)
// 	if err != nil {
// 		return false, err
// 	}

// 	// 3. 判断是否超过阈值
// 	if count >= maxRequests {
// 		return false, nil
// 	}

// 	// 4. 将当前请求添加到 ZSet 中
// 	member := gredis.ZAddMember{
// 		Score:  float64(now),                                 // 分数使用毫秒时间戳
// 		Member: strconv.FormatInt(time.Now().UnixNano(), 10), // 成员使用纳秒时间戳确保唯一
// 	}
// 	_, err = g.Redis().ZAdd(ctx, key, nil, member)
// 	if err != nil {
// 		return false, err
// 	}

// 	// 5. 设置 Key 的过期时间，避免无用数据长期存储
// 	expireTime := windowSize.Seconds()
// 	_, err = g.Redis().Expire(ctx, key, int64(expireTime))
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// // 在中间件中使用
// func SlidingWindowMiddleware(r *ghttp.Request) {
// 	ctx := r.Context()
// 	// key := "sliding_limit:" + r.Router.Uri + ":" + r.GetClientIp()
// 	key := "sliding_limit:" + ":" + r.GetClientIp()
// 	allowed, err := slidingWindowLimit(ctx, key, time.Minute, 100) // 1分钟内最多100次请求
// 	if err != nil {
// 		response.Error(r).SetCode(response.FAILD).
// 			SetMessage("服务内部错误").Send()
// 	}

// 	if !allowed {
// 		response.Error(r).SetCode(response.FAILD).
// 			SetMessage("请求次数超限，请下一个时间窗口再试").Send()
// 	}

// 	r.Middleware.Next()
// }

import (
	"context"
	"server/app/common/utils/response"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// fixedWindowLimit 固定窗口限流函数
// key: 限流键
// windowSize: 窗口大小，如 time.Minute
// maxRequests: 窗口内允许的最大请求数
func fixedWindowLimit(ctx context.Context, key string, windowSize time.Duration, maxRequests int64) (bool, error) {
	// 1. 对 Key 进行计数，如果 Key 不存在则会创建并设置过期时间
	//    使用 INCR 命令原子性增加计数
	count, err := g.Redis().Incr(ctx, key)
	if err != nil {
		return false, err
	}
	// 2. 如果是第一次计数（count == 1），设置键的过期时间
	if count == 1 {

		_, err = g.Redis().Expire(ctx, key, int64(windowSize.Seconds()))
		if err != nil {
			// 如果设置过期时间失败，尝试删除这个 key 避免无过期时间
			g.Redis().Del(ctx, key)
			return false, err
		}
	}

	// 3. 判断计数是否超过限制
	if count > maxRequests {
		return false, nil
	}

	return true, nil
}

func FixedWindowMiddleware(r *ghttp.Request) {
	ctx := r.Context()
	key := "fixed_limit:" + ":" + r.GetClientIp()

	allowed, err := fixedWindowLimit(ctx, key, time.Minute, 200) // 1分钟内最多200次请求
	if err != nil {
		response.Error(r).SetCode(response.FAILD).
			SetMessage("服务内部错误").Send()
	}

	if !allowed {
		response.Error(r).SetCode(response.FAILD).
			SetMessage("请求次数超限，请下一个时间窗口再试").Send()
	}

	r.Middleware.Next()
}

func FixedWindowActionMiddleware(r *ghttp.Request) {
	ctx := r.Context()
	key := "fixed_limit:" + r.Router.Uri + ":" + r.GetClientIp()

	allowed, err := fixedWindowLimit(ctx, key, time.Minute, 5) // 1分钟内最多200次请求
	if err != nil {
		response.Error(r).SetCode(response.FAILD).
			SetMessage("服务内部错误").Send()
	}

	if !allowed {
		response.Error(r).SetCode(response.FAILD).
			SetMessage("请求次数超限，请下一个时间窗口再试").Send()
	}

	r.Middleware.Next()
}
