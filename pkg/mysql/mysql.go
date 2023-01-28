package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Options defines optsions for mysql database.

type Conf struct {
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



var GormPool *gorm.DB
// New create a new gorm db instance with the given options.
func NewMySqlPool(opts *Conf) *gorm.DB{

	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//	Logger: logger.New(&hookSql,opts.LogLevel),
	})
	if err != nil {
		log.Fatalln("init mysql conn poll err",err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("init mysql conn poll err",err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln("init mysql conn poll err",err)
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(opts.MaxConnectionLifeTime) * time.Second)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db
}

func InitMysqlPool(MysqlConf *Conf)  {
	//普通的db方式
	if MysqlConf.Host == "" {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format("2006-01-02 15:04:05"), " empty mysql config.")
	}
	GormPool = NewMySqlPool(MysqlConf)
}

