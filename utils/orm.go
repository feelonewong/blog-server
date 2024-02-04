package utils

import (
	"reflect"
	"strings"
)

func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	if typ.Kind() == reflect.Ptr {
		//	解析指针类型, 指针转换为非指针
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			if fieldType.IsExported() {
				//	只关注可导出成员
				if fieldType.Tag.Get("gorm") == "-" {
					// -表示go里面的结构体不对应MySQL的某一列
					continue
				}
				name := Camel2Snake(fieldType.Name) // 没有gorm tag 则将驼峰转为蛇形
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						pos := strings.Index(content, ";")
						if pos > 0 {
							name = content[0:pos]
						} else if pos < 0 {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	} else {
		return nil
	}
}
