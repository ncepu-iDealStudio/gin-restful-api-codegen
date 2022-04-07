// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 17:52
// @Software: GoLand

package utils

import "reflect"

//binding type interface 要修改的结构体
//value type interface 有数据的结构体

func StructAssign(binding interface{}, value interface{}, tagName string) {
	bVal := reflect.ValueOf(binding)
	if bVal.Kind() == reflect.Ptr {
		bVal = bVal.Elem()
	} //获取reflect.Type类型
	vVal := reflect.ValueOf(value)
	if vVal.Kind() == reflect.Ptr {
		vVal = vVal.Elem()
	} //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	bTypeOfT := bVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		if !vVal.Field(i).IsZero() {
			vName := vTypeOfT.Field(i).Tag.Get(tagName)
			for j := 0; j < bVal.NumField(); j++ {
				bName := bTypeOfT.Field(j).Tag.Get(tagName)
				if vName == bName {
					bVal.Field(j).Set(reflect.ValueOf(vVal.Field(i).Interface()))
					break
				}
			}
		}
	}
}
