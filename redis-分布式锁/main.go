package redis

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

type RedisLock struct {
	key      string
	value    string
	duration time.Duration
}

// 随机字符串 保证不重复
func getRandString() string {
	return strings.ToLower(strconv.FormatInt(time.Now().UnixNano(), 36))
}

// 默认过期时间
var DEFAULT_EXPIRE_TIME = 100 * time.Millisecond

// 新建
func NewRedisLock(key string, expire ...time.Duration) RedisLock {
	var duration time.Duration
	if len(expire) == 0 {
		duration = DEFAULT_EXPIRE_TIME
	} else {
		duration = expire[0]
	}
	return RedisLock{
		key:      "Mutexkey:" + key,
		value:    getRandString(),
		duration: duration,
	}
}

// 加锁
func (r *RedisLock) KeyLock(ctx context.Context, redisLock RedisLock) (bool, error) {
	setLock, err := rdb.SetNX(ctx, redisLock.key, redisLock.value, redisLock.duration).Result()
	if err != nil {
		return false, err
	}
	return setLock, nil
}

// 释放锁
func (r *RedisLock) DelByKeyWhenValueEquals(ctx context.Context, rdb *redis.Client, redisLock RedisLock) (bool, error) {
	lua := `
        if redis.call('GET', KEYS[1]) == ARGV[1] then
         	return redis.call('DEL', KEYS[1])
        else
	       return 0
        end
        `
	scriptKeys := []string{redisLock.key}
	val, err := rdb.Eval(ctx, lua, scriptKeys, redisLock.value).Result()
	if err != nil {
		return false, err
	}
	return val == int64(1), nil
}

// 守护线程 延迟锁的过期时间
func (r *RedisLock) WatchDog(ctx context.Context, rdb *redis.Client, redisLock RedisLock) {
	for {
		select {
		// 业务完成
		case <-ctx.Done():
			fmt.Printf("任务完成,关闭%s的自动续期\n", redisLock.key)
			return
		// 业务未完成
		default:
			// 自动续期
			rdb.PExpire(ctx, redisLock.key, redisLock.duration)
			// 继续等待
			time.Sleep(redisLock.duration / 2)
		}
	}
}
