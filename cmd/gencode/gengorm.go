package main

import (
	"github.com/qmhball/db2gorm/gen"
)

func main() {
	dsn := "root:@tcp(192.168.1.13:3306)/igoal_sit2?charset=utf8&parseTime=true&loc=Local"
	//生成指定单表
	tblName := "t_user"
	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}
