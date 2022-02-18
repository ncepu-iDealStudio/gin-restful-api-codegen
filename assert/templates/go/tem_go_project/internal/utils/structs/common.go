// coding: utf-8
// @Author : lryself
// @Date : 2021/5/16 2:40
// @Software: GoLand

package structs

import (
	"fmt"
	"reflect"
)

func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			temp := v.Field(i).Interface()
			if !reflect.DeepEqual(temp, reflect.Zero(reflect.TypeOf(temp)).Interface()) {
				out[tagValue] = temp
			}
		}
	}
	return out, nil
}

func MapToStruct(in map[string]interface{}, out interface{}) error {
	v := reflect.ValueOf(out)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for key, value := range in {
		_, ok := t.FieldByName(key)
		if ok {
			v.FieldByName(key).Set(reflect.ValueOf(value))
		}
	}

	return nil
}
