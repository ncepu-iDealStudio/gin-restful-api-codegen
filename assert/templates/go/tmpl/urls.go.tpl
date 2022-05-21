// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package api

import (
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/apis/api_1_0"{{range $CodeDict.TablesInfo}}
    "{{$CodeDict.Dict.ProjectName}}/internal/routers/api/{{.PackageName}}"
{{end}}
)

var (
    Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
    Api = engine.Group("api")
    Api.Any("version", api_1_0.GetVersion)
{{range $CodeDict.TablesInfo}}
    {{.PackageName}}.Init{{.StructName}}RouterGroup(Api)
{{end}}
}