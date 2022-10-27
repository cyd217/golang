package pool

import (
	"context"
	"delay_demo/pack"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	redis2 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	JobPoolKey         = "job_pool_key_"
	BaseDelayBucketKey = "base_delay_bucket"
	BaseReadyQueueKey  = "base_ready_queue"
)

type BodyContent struct {
	OrderID   int
	OrderName string
}
type RedisJobData struct {
	Topic string
	ID    string
	Delay int
	TTL   int
	Body  *BodyContent
}

func (d *RedisJobData) SetJobPool(ctx context.Context, number int) bool {
	redisCoon := pack.GetRedisDb()

	for i := 0; i < number; i++ {
		d.Topic = "order_queue"
		d.ID = uuid.NewString()
		d.Delay = 1
		d.TTL = 3
		d.Body = &BodyContent{
			OrderID:   i,
			OrderName: "order_name_" + strconv.Itoa(i),
		}
		key := JobPoolKey + strconv.Itoa(i)
		data, _ := json.Marshal(d)
		//写入job pool
		_, err := redisCoon.Set(ctx, key, data, 0*time.Second).Result()
		if err != nil {
			fmt.Println("添加失败: ", err)
			return false
		}
		fmt.Println("添加成功: ", key)
		nowTime := time.Now().Unix()
		rand.Seed(time.Now().UnixNano())
		delayTime := int(nowTime) + rand.Intn(30)
		//写入delay queue
		_, err = redisCoon.ZAdd(ctx, BaseDelayBucketKey, &redis2.Z{
			Score:  float64(delayTime),
			Member: key,
		}).Result()
		if err != nil {
			fmt.Println("添加失败:zadd ", err)
			return false
		}
		fmt.Println("添加成功: zadd", BaseDelayBucketKey)
	}

	return false
}

//定时timer.tick查询bucket中是否有过期的数据，如果有放入消费队列中
func TimerDelayBucket(redisCoon *redis2.Client, ctx context.Context, pool *Pool) error {
	nowTime := time.Now().Unix()
	result, err := redisCoon.ZRangeByScoreWithScores(ctx, BaseDelayBucketKey, &redis2.ZRangeBy{
		Min: "-inf",
		Max: strconv.FormatInt(nowTime, 10),
	}).Result()
	if err == nil {
		for _, z := range result {
			//进入ready queue
			redisCoon.LPush(ctx, BaseReadyQueueKey, z.Member)
			//写入通道说明有数据了，可以进行消费
			err := pool.Put(&Task{
				Member: z.Member.(string),
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return err

}

//消费队列
func ConsumeQueue(ctx context.Context, redisCoon *redis2.Client) error {
	//先判断list中是否有数据，有数据才有必要执行，没有数据直接返回就好了
	lenQueue, err := redisCoon.LLen(ctx, BaseReadyQueueKey).Result()
	if err != nil {
		return err
	}
	if lenQueue == 0 {
		return nil
	}
	result, err := redisCoon.LPop(ctx, BaseReadyQueueKey).Result()
	if err != nil {
		return err
	}
	fmt.Println("我消费了一个数据：", result)
	//这里可以实现需要的操作，这里简单实现了删除操作
	redisCoon.Del(ctx, result)
	redisCoon.ZRem(ctx, BaseDelayBucketKey, result)
	return nil
}
