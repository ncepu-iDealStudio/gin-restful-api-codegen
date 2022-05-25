// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package models

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type {{$CodeDict.TableInfo.StructName}}Model struct { {{range .TableInfo.Columns}}
    {{.Field}} {{.DataType}} `gorm:"column:{{.Field}};{{.TagKey}}{{.TagColumnType}}{{.TagNull}}{{.TagDefault}}" json:"{{.Field}}" form:"{{.Field}}"`{{end}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) TableName() string {
    return "{{$CodeDict.TableInfo.TableName}}"
}
{{range .TableInfo.Columns}}
func (m *{{$CodeDict.TableInfo.StructName}}Model) Get{{.Field}}() {{.DataType}} {
    return m.{{.Field}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Set{{.Field}}({{.Field}} {{.DataType}}) {
    m.{{.Field}} = {{.Field}}
}
{{end}}

func (m *{{$CodeDict.TableInfo.StructName}}Model) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Assign(in interface{}) {
    structs.StructAssign(m, in, "json")
}