package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type MyStruct struct {
	Id         int    `gorm:"column:id;primaryKey"`
	PassWd     string `json:"passWd" gorm:"column:password"`
	Name       string
	Family     string    `gorm:"-"`
	CreateTime time.Time `form:"create_time" binding:"required,before_today" time_formate:"2006-01-02" time_utc:"8"`
}

func PrintFieldInfo(object any) {
	tp := reflect.TypeOf(object)
	fieldNum := tp.NumField()
	for i := 0; i < fieldNum; i++ {
		field := tp.Field(i)
		fmt.Printf("%d %s offset:%d anonymous:%t type:%s exported:%t gorm tag=%s json tag:%s\n", i,
			field.Name,
			field.Offset,          // 偏移量
			field.Anonymous,       // 是否匿名
			field.Type,            // 字段类型
			field.IsExported(),    // 是否支持导出
			field.Tag.Get("gorm"), // 获取gorm的字段
			field.Tag.Get("json"), // 获取json的字段
		)
	}
	fmt.Println()
}

func TestPrintFieldInfo(t *testing.T) {
	PrintFieldInfo(MyStruct{})
}
