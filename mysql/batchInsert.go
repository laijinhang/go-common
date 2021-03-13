package mysql

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// 不支持结构体内嵌结构体
// BatchInsert 批量插入数据
// sql := "insert into `table_name` (...) values"
func BatchInsert(tableName string, data interface{}) string {
	// 1、如果data不是切片
	if !isSlice(data) {
		panic(fmt.Sprintf("%s not slice", data))
	}
	// 2、如果切片长度等于0
	vap := reflect.ValueOf(data)
	if vap.Len() == 0 {
		panic("")
	}

	// 3、获取表字段名
	tagp := vap.Index(0).Type().Elem()
	columns := make([]string, 0)

	for i := 0; i < tagp.NumField(); i++ {
		columns = append(columns, func() string {
			tag := fmt.Sprintf(tagp.Field(i).Tag.Get("gorm"))
			tag = strings.Split(tag, ",")[0]
			return "`" + strings.Split(tag, ":")[1] + "`"
		}())
	}
	var buffer bytes.Buffer
	// insert into `table_name`
	if _, err := buffer.WriteString(fmt.Sprintf("insert into `%s`", tableName)); err != nil {
		panic(err)
	}
	buffer.WriteString(fmt.Sprintf(" (%s) VALUES ", strings.Join(columns, ",")))

	// 4、获取表值
	for i := 0; i < vap.Len(); i++ {
		buffer.WriteString("(")
		for j := 0; j < vap.Index(i).Elem().NumField(); j++ {
			buffer.WriteString(fmt.Sprintf("'%s'", fmt.Sprint(vap.Index(i).Elem().Field(j))))
			if j != vap.Index(i).Elem().NumField()-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString(")")
		if i == vap.Len()-1 {
			buffer.WriteString(";")
		} else {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func isSlice(data interface{}) bool {
	typ := reflect.ValueOf(data)
	t := typ.Type().String()
	return len(t) > 2 && t[0] == '[' && t[1] == ']'
}
