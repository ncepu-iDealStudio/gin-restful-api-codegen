// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "time"
)
{{ $StructName := .TableInfo.StructName }}
type {{$StructName}} struct {
    {{range .TableInfo.Columns}}
    {{.Field}} {{.GoType}} `gorm:"column:{{.Field}};{{.Key}}{{.Type}}{{.Null}}" json:"{{.Field}}" form:"{{.Field}}"`{{end}}
}

func (m *{{$StructName}}) TableName() string {
    return "approval_collegeTask"
}

{{range .TableInfo.Columns}}
func (m *{{$StructName}}) Get{{.Field}}() {{.GoType}} {
    return m.{{.Field}}
}

func (m *{{$StructName}}) Set{{.Field}}({{.Field}} {{.GoType}}) {
    m.{{.Field}} = {{.Field}}
}
{{end}}