package unit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
)

func VerifyPhone(phone string) bool {
	return regexp.MustCompile( `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`).MatchString(phone)
}

func StructToMap(i interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	typeC := reflect.TypeOf(i)
	valC := reflect.ValueOf(i)

	fieldNum := typeC.NumField()
	for i := 0;i < fieldNum;i++ {
		switch valC.Field(i).Type().String() {
		case "int64":
			m[typeC.Field(i).Name] = valC.Field(i).Int()
		case "string":
			m[typeC.Field(i).Name] = valC.Field(i).String()
		case "interface {}":
			m[typeC.Field(i).Name] = valC.Field(i).Interface()
		}
	}
	return m
}

func ToJsonString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	return string(bytes), err
}
// json: unsupported type: map[interface {}]interface {}

func ToTypeString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func MapInterfaceToString(m map[interface{}]interface{}) map[interface{}]string {
	ms := make(map[interface{}]string)

	for k, v := range m {
		ms[k] = fmt.Sprintf("%v", v)
		/* switch reflect.TypeOf(v).Name() {
		   case "int":
		       ms[k] = return fmt.Sprintf("%v", v)
		   case "int64":
		       ms[k] = strconv.FormatInt(v.(int64), 10)
		   case "string":
		       ms[k] = v.(string)
		   } */
	}
	return ms
}
