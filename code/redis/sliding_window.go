package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

func isActionAllowdByZset(client *redis.Client, userID, actionKey string, period time.Duration, maxCount int64) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("hist:%s:%s", userID, actionKey)
	mperiod := period.Nanoseconds() / 1e6       // 转毫秒
	now := int64(time.Now().Nanosecond() / 1e6) // 毫秒时间戳
	// 注意这里不能使用 now = time.Now().Seconds()*1000 ，这样做的话会丢失精度，导致一秒内的所有now值都一样;

	pipe := client.Pipeline()
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now), //记录行为，score使用毫秒时间戳;
		Member: uuid.NewV4(),
	})
	// 移除时间窗口之前的行为记录, 剩下的都是时间窗口内的
	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%v", now-mperiod))
	// 获取窗口内的行为数量
	pipe.ZCard(ctx, key)
	// 设置zset 过期时间, 避免冷数据持续占用内存
	// 过期时间应该等于时间窗口的长度, 再多宽限1s
	pipe.Expire(ctx, key, period+1)
	// 执行
	res, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	cmd, ok := res[2].(*redis.IntCmd)
	if ok {
		return cmd.Val() <= maxCount, nil
	}
	return false, errors.New(fmt.Sprintf("parse redis result failed, origin result: %v", res))
}

func isActionAllowdByHash(client *redis.Client, userID, actionKey string, period time.Duration, maxCount int64) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("hist:%s:%s:%s", userID, actionKey, period)

	timeLayout := "20060102150405"
	now := time.Now()
	nowStr := now.Format(timeLayout)
	expiredTime := now.Add(-period)

	// 增加一次访问记录
	client.HIncrBy(ctx, key, nowStr, 1)
	// 设置过期时间=窗口时间+1
	client.Expire(ctx, key, period+1)
	// 取出所有的计数器
	all := client.HGetAll(ctx, key)
	// 移除时间窗口之前的行为记录, 剩下的都是时间窗口内的
	var sum int64 = 0
	onDeleteFields := make([]string, 0)
	for k, v := range all.Val() {
		originTime, err := time.Parse(timeLayout, k)
		if err != nil {
			return false, err
		}
		if originTime.Before(expiredTime) {
			onDeleteFields = append(onDeleteFields, k)
		} else {
			vv, _ := strconv.ParseInt(v, 10, 64)
			sum += vv
		}
	}
	if len(onDeleteFields) > 0 {
		client.HDel(ctx, key, onDeleteFields...)
	}
	return sum <= maxCount, nil
}
