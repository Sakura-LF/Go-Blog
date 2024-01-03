package init

import (
	"Go-Blog/util"
	"errors"
	"github.com/spf13/viper"
)

var (
	LogConfig   *viper.Viper
	MysqlConfig *viper.Viper
)

func Init() {
	LogConfig = LoadConfig("log")
	MysqlConfig = LoadConfig("mysql")
}

func LoadConfig(filename string) *viper.Viper {
	config := viper.New()
	rootPath := util.GetRootPath()
	config.AddConfigPath(rootPath + "/config")
	config.SetConfigName(filename)
	err := config.ReadInConfig()
	if err != nil {
		// 如果需要对配置文件不存在错误，做特殊处理，使用：
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			//config file not found; ignore error if desired
			panic("config file was not found: " + filename)
		} else {
			panic(err)
		}
	}

	return config
}
