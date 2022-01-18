// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"errors"
	"fmt"
	"forest/log/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Options defines optsions for mysql database.

// New create a new gorm db instance with the given options.
func NewMySqlPool(opts *MySQLConf) (*gorm.DB, error) {
	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.New(opts.LogLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(opts.MaxConnectionLifeTime) * time.Second)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}

func InitMysqlPool(MysqlConfMap *MysqlMapConf) error {
	//普通的db方式
	if len(MysqlConfMap.Mysql) == 0 {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format("2006-01-02 15:04:05"), " empty mysql config.")
	}

	GormMapPool = map[string]*gorm.DB{}
	for confName, DbConf := range MysqlConfMap.Mysql {
		dbgorm, err := NewMySqlPool(DbConf)
		if err != nil {
			return err
		}
		GormMapPool[confName] = dbgorm
	}

	return nil
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbpool, ok := GormMapPool[name]; ok {
		return dbpool, nil
	}
	return nil, errors.New("get mysql pool error")
}
