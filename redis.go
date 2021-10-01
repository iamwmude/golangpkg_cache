package cache

import (
	"context"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/iamwmude/dictionary/db/pkg/utils"

	jsoniter "github.com/json-iterator/go"
)

type Redis struct {
	rdb *goredislib.Client
	ctx context.Context
	rs  *redsync.Redsync
}

func (cache *Redis) Set(key string, value interface{}, timeout time.Duration) error {
	var setVal interface{}
	switch value.(type) {
	case string, bool, float32, float64, int, int8, int16, int32, int64:
		setVal = value
	default:
		setVal = utils.GetString(value)
	}

	return cache.rdb.Set(cache.ctx, key, setVal, timeout).Err()
}

func (cache *Redis) Delete(key string) error {
	return cache.rdb.Del(cache.ctx, key).Err()
}

func (cache *Redis) GetString(key string) (string, error) {
	return cache.rdb.Get(cache.ctx, key).Result()
}

func (cache *Redis) GetBool(key string) (bool, error) {
	return cache.rdb.Get(cache.ctx, key).Bool()
}

func (cache *Redis) GetFloat64(key string) (float64, error) {
	return cache.rdb.Get(cache.ctx, key).Float64()
}

func (cache *Redis) GetInt(key string) (int, error) {
	return cache.rdb.Get(cache.ctx, key).Int()
}

func (cache *Redis) GetInt64(key string) (int64, error) {
	return cache.rdb.Get(cache.ctx, key).Int64()
}

func (cache *Redis) GetStructData(key string, data interface{}) error {
	val, err := cache.rdb.Get(cache.ctx, key).Result()
	if nil != err {
		return err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(([]byte)(val), data)
	if nil != err {
		return err
	}

	return nil
}

func (cache *Redis) GetMutex(key string) *redsync.Mutex {
	return cache.rs.NewMutex(key)
}

func (cache *Redis) Increase(key string) (int64, error) {
	return cache.rdb.Incr(cache.ctx, key).Result()
}

func (cache *Redis) Decrease(key string) (int64, error) {
	return cache.rdb.Decr(cache.ctx, key).Result()
}
