package utils

import (
	"github.com/bytedance/sonic"
	"math/rand"
)

type User struct {
	Name string
}

// struct转换为JSON类型
func Rs() string {
	u := User{Name: "黄飞龙"}
	brr, _ := sonic.Marshal(u)
	return string(brr)
}

// 判断是否为大写
func IsASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

// 大小写转换
func UpperLowerExchange(c byte) byte {
	return c ^ ' '
}

// 驼峰转换为蛇形 UserName user_name
func Camel2Snake(s string) string {
	if len(s) == 0 {
		return ""
	}
	t := make([]byte, 0, len(s)+4)
	if IsASCIIUpper(s[0]) {
		t = append(t, UpperLowerExchange(s[0]))
	} else {
		t = append(t, s[0])
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if IsASCIIUpper(c) {
			t = append(t, '_', UpperLowerExchange(c))
		} else {
			t = append(t, c)
		}
	}
	return string(t)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 生成随机字符串
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
