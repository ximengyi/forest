package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)


type Conf struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

var redisConf *Conf
var RedisMapPool map[int]*redis.Client

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
	redisConf = &Conf{}
	redisConf.Addr = addr
	redisConf.Password = passwd
	_,err := RedisConnFactory(0)
	return err
}