// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package mysqlModel

import (
    "{{$CodeDict.Dict.ProjectName}}/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type {{$CodeDict.TableInfo.StructName}}Model struct { {{range .TableInfo.Columns}}
    {{.Field}} {{.GoType}} `gorm:"column:{{.Field}};{{.TagKey}}{{.TagColumnType}}{{.TagNull}}{{.TagDefault}}" json:"{{.Field}}" form:"{{.Field}}"`{{end}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) TableName() string {
    return "{{$CodeDict.TableInfo.TableName}}"
}
{{range .TableInfo.Columns}}
func (m *{{$CodeDict.TableInfo.StructName}}Model) Get{{.Field}}() {{.GoType}} {
    return m.{{.Field}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Set{{.Field}}({{.Field}} {{.GoType}}) {
    m.{{.Field}} = {{.Field}}
}
{{end}}

func (m *{{$CodeDict.TableInfo.StructName}}Model) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}