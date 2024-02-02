package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	LogRus *logrus.Logger
)

func InitLog(configFile string) {
	viper := CreateConfig(configFile)
	LogRus = logrus.New()
	switch strings.ToLower(viper.GetString("level")) {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", viper.GetString("level")))
	}
	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	// 全局路径 + 获取配置文件的config.yaml/file属性
	logFile := ProjectRootPath + "/" + viper.GetString("file")
	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H",                      // 指定日志文件的路径和名称，路径不存在的时候会创建
		rotatelogs.WithLinkName(logFile),         // 为最新的一份日志创建软连接
		rotatelogs.WithRotationTime(1*time.Hour), // 每隔一小时生成一份新的文件
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 只保留最近7天的日志,
	)
	if err != nil {
		panic(err)
	}
	LogRus.SetOutput(fout)       // 设置输出日志文件
	LogRus.SetReportCaller(true) // 输出从哪里调起的日志打印
}
