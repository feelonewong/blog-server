package test

import (
	"blog-server/utils"
	"fmt"
	"path"
	"runtime"
	"testing"
	"time"
)

func af() (string, int) {
	// 0 1 2
	_, filename, line, _ := runtime.Caller(2)
	return filename, line
}

func bf() (string, int) {
	return af()
}

// TestCaller -> bf -> af  形参0调用第1个(定义方法的地方)/1调用第2个/2调用第3个(最终调用的地方)
func TestCaller(t *testing.T) {
	filename, line := bf()
	fmt.Println(path.Dir(filename), line)
}

func TestConfig(t *testing.T) {
	dbViper := utils.CreateConfig("mysql")
	// 确保调用前添加了配置路径.实时监听配置文件变化
	dbViper.WatchConfig()
	// 读取配置的第一种方式
	if !dbViper.IsSet("blog.port") {
		t.Fail() // 检查有没有此项配置
	}
	port := dbViper.GetInt("blog.port")
	fmt.Println("端口号port:", port)
	time.Sleep(10 * time.Second) // 休息10秒钟
	port2 := dbViper.GetInt("blog.port")
	fmt.Println("端口号2port:", port2)
}
