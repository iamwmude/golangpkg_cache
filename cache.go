package cache

import (
	"context"
	"log"

	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type Cache interface {
	Set(key string, value interface{}, timeout time.Duration) error
	Delete(key string) error
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetFloat64(key string) (float64, error)
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetStructData(key string, data interface{}) error
	Increase(key string) (int64, error)
	Decrease(key string) (int64, error)
}

var cache Cache

func InitRedisCache(options goredislib.Options) {
	redis := Redis{
		ctx: context.Background(),
	}

	if redis.rdb = goredislib.NewClient(&options); redis.rdb == nil {
		log.Fatal("init cache failed")
	} else {
		if _, err := redis.rdb.Ping(redis.ctx).Result(); err != nil {
			log.Fatal("init cache failed")
		}
	}

	pool := goredis.NewPool(redis.rdb)
	if redis.rs = redsync.New(pool); redis.rs == nil {
		log.Fatalf("init redsync failed")
	}

	cache = &redis
}
