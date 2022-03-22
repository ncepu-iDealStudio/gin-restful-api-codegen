// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package api_1_0

import (
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/apis/api1_0"{{range $CodeDict.TablesInfo}}
    "{{$CodeDict.Dict.ProjectName}}/internal/routers/api1_0/{{.PackageName}}"
{{end}}
)

var (
    Api *gin.RouterGroup
)

func InitAPI_1_0Router(engine *gin.Engine) {
    Api = engine.Group("api_1_0")
    Api.Any("version", api_1_0.GetVersion)
{{range $CodeDict.TablesInfo}}
    {{.PackageName}}.Init{{.StructName}}RouterGroup(Api)
{{end}}
}