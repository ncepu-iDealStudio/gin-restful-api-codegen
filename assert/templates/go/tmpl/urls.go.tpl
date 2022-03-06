// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package api1_0

import (
    "github.com/gin-gonic/gin"
    "{{$CodeDict.ProjectName}}/internal/apis/api1_0"{{range $CodeDict.Tables}}
    "{{$CodeDict.ProjectName}}/internal/routers/api1_0/{{.PackageName}}"
{{end}}
)

var (
    Api *gin.RouterGroup
)

func InitAPI1_0Router(engine *gin.Engine) {
    Api = engine.Group("api1_0")
    Api.Any("version", api1_0.GetVersion)
{{range .Tables}}
    {{.PackageName}}.Init{{.StructName}}RouterGroup(Api)
{{end}}
}