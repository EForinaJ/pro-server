package utlis_lock

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func SetCount(ctx context.Context, key, lockKey string, expired int64, count int) (int, error) {
	Times := 0
	TimeObj, err := g.Redis().Get(ctx, key)
	if err != nil {
		return Times, err
	}
	if !TimeObj.IsEmpty() {
		TimeIncr, err := g.Redis().Incr(ctx, key)
		Times = gconv.Int(TimeIncr)
		if err != nil {
			return Times, err
		}
	} else {
		err := g.Redis().SetEX(ctx, key, Times, expired)
		if err != nil {
			return Times, err
		}
	}

	if Times >= count {
		setLock(ctx, lockKey, expired)
	}
	return Times, nil
}

func setLock(ctx context.Context, key string, expired int64) {
	_ = g.Redis().SetEX(ctx, key, "true", expired)
}

// CheckLock 检查账号是否锁定
func CheckLock(ctx context.Context, key string) bool {
	result := false
	lock, err := g.Redis().Get(ctx, key)
	if err != nil {
		return false
	}
	if lock != nil {
		lockContent := gconv.String(lock)
		if lockContent == "true" {
			result = true
		}
	}
	return result
}
