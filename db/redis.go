package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisConf *RedisConf

func RedisConnFactory(db int) (*redis.Client, error) {
	if RedisMapPool[db] != nil {
		return RedisMapPool[db], nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password, // no password set
		DB:       db,                                      // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	RedisMapPool[db] = rdb
	return RedisMapPool[db], nil
}

func InitRedisConf(addr string,passwd string) error {
	if RedisMapPool == nil {
		RedisMapPool = make(map[int]*redis.Client)
	}
	redisConf = &RedisConf{}
	redisConf.Addr = addr
	redisConf.Password = passwd
	_,err := RedisConnFactory(0)
	return err
}