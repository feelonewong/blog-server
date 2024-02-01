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

//go test -v ./utils/test -run=^TestCamel2Snake$ -count=1
