package db

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type MysqlMapConf struct {
	Mysql map[string]*MySQLConf `mapstructure:"mysql"`
}

type MySQLConf struct {
	Host                  string `mapstructure:"host"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	Database              string `mapstructure:"database"`
	DriverName            string `mapstructure:"driver_name"`
	MaxOpenConnections    int    `mapstructure:"max_open_conn"`
	MaxIdleConnections    int    `mapstructure:"max_idle_conn"`
	MaxConnectionLifeTime int    `mapstructure:"max_conn_life_time"`
	LogLevel              int    `mapstructure:"log_level"`
}

type RedisConf struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type MongodbConf struct {
	Host                  string        `mapstructure:"host"`
	Username              string        `mapstructure:"username"`
	Password              string        `mapstructure:"password"`
	Database              string        `mapstructure:"database"`
	MaxOpenConnections    uint64        `mapstructure:"max_open_conn"`
	MaxPoolSize           uint64        `mapstructure:"max_pool_size"`
	MaxConnectionIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
	LogLevel              int           `mapstructure:"log_level"`
}

// 全局变量
//var ConfBase *BaseConf
//var DBMapPool map[string]*sql.DB
var GormMapPool map[string]*gorm.DB

//var GORMDefaultPool *gorm.DB

var Mongodb *mongo.Database
var MongodbClient *mongo.Client

//var ConfRedis *RedisConf
//var ViperConfMap map[string]*viper.Viper
var RedisMapPool map[int]*redis.Client

// 获取基本配置信息
