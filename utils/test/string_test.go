package test

import (
	"blog-server/utils"
	"fmt"
	"testing"
)

func TestCamel2Snake(t *testing.T) {
	s1 := "Abc"
	s2 := utils.Camel2Snake(s1)
	if s2 != "abc" {
		fmt.Println(s2)
		t.Fail()
	}
	s1 = "AbcEfg"
	s2 = utils.Camel2Snake(s1)
	if s2 != "abc_efg" {
		fmt.Println(s2)
		t.Fail()
	}
	s1 = "abcEfg"
	s2 = utils.Camel2Snake(s1)
	if s2 != "abc_efg" {
		fmt.Println(s2)
		t.Fail()
	}
}

// 性能基准测试
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.Camel2Snake("UserName")
	}
}

// 生成随机数
func TestRandStringRunes(t *testing.T) {
	fmt.Println(utils.RandStringRunes(20))
}

//go test -v ./utils/test -run=^TestCamel2Snake$ -count=1
