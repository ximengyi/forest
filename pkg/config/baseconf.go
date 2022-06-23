package config

type BaseConf struct {
	DebugMode    string    `mapstructure:"debug_mode"`
	TimeLocation string    `mapstructure:"time_location"`
	Log          LogConfig `mapstructure:"log"`
	JWTConfig    JWTConf   `mapstructure:"jwt"`
	OSSConfig    OSSConf   `mapstructure:"oss"`
	RedisConfig  RedisConf `mapstructure:"redis"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
}

type JWTConf struct {
	Key      string `mapstructure:"key"`
	Duration int64  `mapstructure:"duration"`
}

type OSSConf struct {
	Key      string `mapstructure:"key"`
	Secret   string `mapstructure:"secret"`
	Endpoint string `mapstructure:"endpoint"`
	Domain   string `mapstructure:"domain"`
}

type LogConfFileWriter struct {
	On              bool   `mapstructure:"on"`
	LogPath         string `mapstructure:"log_path"`
	RotateLogPath   string `mapstructure:"rotate_log_path"`
	WfLogPath       string `mapstructure:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_log_path"`
}

type LogConfConsoleWriter struct {
	On    bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type LogConfig struct {
	Level              string   `mapstructure:"level"`
	Format             string   `mapstructure:"format"`
	EnableColor        bool     `mapstructure:"enable_color"`
	DisableCaller      bool     `mapstructure:"disable_caller"`
	Development        bool     `mapstructure:"development"`
	DisableStacktrace  bool     `mapstructure:"disable_stacktrace"`
	OutputPaths        []string `mapstructure:"output_paths"`
	ErrorOutputPaths   []string `mapstructure:"error_output_paths"`
}


type RedisConf struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// 全局变量
var ConfBase *BaseConf
//var DBMapPool map[string]*sql.DB
//var GORMMapPool map[string]*gorm.DB
//var DBDefaultPool *sql.DB
//var GORMDefaultPool *gorm.DB
//var ConfRedis *RedisConf
//var ViperConfMap map[string]*viper.Viper
//var RedisMapPool map[int]*redis.Client
// 获取基本配置信息
func GetBaseConf() *BaseConf {
	return ConfBase
}

func InitBaseConf(path string) error {

	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	//if ConfBase.DebugMode == "" {
	//	if ConfBase.Base.DebugMode != "" {
	//		ConfBase.DebugMode = ConfBase.Base.DebugMode
	//	} else {
	//		ConfBase.DebugMode = "debug"
	//	}
	//}
	//if ConfBase.TimeLocation == "" {
	//	if ConfBase.Base.TimeLocation != "" {
	//		ConfBase.TimeLocation = ConfBase.Base.TimeLocation
	//	} else {
	//		ConfBase.TimeLocation = "Asia/Shanghai"
	//	}
	//}
	//if ConfBase.Log.Level == "" {
	//	ConfBase.Log.Level = "trace"
	//}

	return nil
}




