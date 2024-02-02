package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var (
	ProjectRootPath = path.Dir(getOncurrentPath() + "/../")
)

func getOncurrentPath() string {
	// 0 表示当前代码在什么位置
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

// Viper解析JSON/TOML/YAML/HCL/INT/ENV等格式配置文件，不需要重启就可以读到最新的值
// 约定配置文件都存储在config/目录下面，格式都是yaml格式
func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := ProjectRootPath + "/config/"
	config.AddConfigPath(configPath) // 文件所在目录
	config.SetConfigName(file)       // 文件名
	config.SetConfigType("yaml")     // 设置文件类型
	configFile := configPath + file + ".yaml"

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 系统初始化阶段发生任何错误，直接结束进程
			panic(fmt.Errorf("找不到配置文件: %s", configFile))
		} else {
			panic(fmt.Errorf("解析配置文件%s出错: %s", configFile, err))
		}
	}
	return config
}
