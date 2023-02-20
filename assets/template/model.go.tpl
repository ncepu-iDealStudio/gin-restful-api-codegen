// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package models

import (
    "{{$CodeDict.Dict.ProjectName}}/utils/structs"
    {{if .TableInfo.HasTimeField}} "time" {{end}}
)

type {{$CodeDict.TableInfo.StructName}}Model struct { {{range .TableInfo.Columns}}
    {{.FieldName}} {{.DataType}} `gorm:"column:{{.Field}};{{.TagKey}}{{.TagColumnType}}{{.TagNull}}{{.TagDefault}}" json:"{{.Field}}" form:"{{.Field}}"`{{end}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) TableName() string {
    return "{{$CodeDict.TableInfo.TableName}}"
}
{{range .TableInfo.Columns}}
func (m *{{$CodeDict.TableInfo.StructName}}Model) Get{{.FieldName}}() {{.DataType}} {
    return m.{{.FieldName}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Set{{.FieldName}}({{.FieldName}} {{.DataType}}) {
    m.{{.FieldName}} = {{.FieldName}}
}
{{end}}

func (m *{{$CodeDict.TableInfo.StructName}}Model) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Assign(in interface{}) {
    structs.StructAssign(m, in, "json")
}