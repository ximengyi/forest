package config

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)


var ConfEnvPath ="./env/"


func GetConfPath(fileName string) string {
	return ConfEnvPath + "/" + fileName + ".toml"
}

func GetConfFilePath(fileName string) string {
	return ConfEnvPath + "/" + fileName
}


//本地解析文件
func ParseLocalConfig(fileName string, st interface{}) error {
	path := GetConfFilePath(fileName)
	err := ParseConfig(path, st)
	if err != nil {
		return err
	}
	return nil
}

//初始化配置文件
func InitViperConf() error {
	f, err := os.Open(ConfEnvPath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, errRead := ioutil.ReadFile(ConfEnvPath + "/" + f0.Name())
			if errRead != nil {
				return errRead
			}
			v := viper.New()
			v.SetConfigType("toml")
			errRead = v.ReadConfig(bytes.NewBuffer(bts))
			if errRead != nil {
				return errRead
			}

			pathArr := strings.Split(f0.Name(), ".")
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}

	return nil
}
