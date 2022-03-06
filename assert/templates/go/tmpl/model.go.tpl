// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package mysqlModel

import (
    "time"
)

type {{$CodeDict.TableInfo.StructName}}Model struct { {{range .TableInfo.Columns}}
    {{.Field}} {{.GoType}} `gorm:"column:{{.Field}};{{.TagKey}}{{.TagColumnType}}{{.TagNull}}{{.TagDefault}}" json:"{{.Field}}" form:"{{.Field}}"`{{end}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) TableName() string {
    return "approval_collegeTask"
}
{{range .TableInfo.Columns}}
func (m *{{$CodeDict.TableInfo.StructName}}Model) Get{{.Field}}() {{.GoType}} {
    return m.{{.Field}}
}

func (m *{{$CodeDict.TableInfo.StructName}}Model) Set{{.Field}}({{.Field}} {{.GoType}}) {
    m.{{.Field}} = {{.Field}}
}
{{end}}