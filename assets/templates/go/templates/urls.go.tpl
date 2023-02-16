// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package api_1_0

import (
    "github.com/gin-gonic/gin"{{range $CodeDict.TablesInfo}}
    "{{$CodeDict.Dict.ProjectName}}/internal/apis/api_1_0/{{.PackageName}}"{{end}}
)

var (
    Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
    Api = engine.Group("api")
    Api.Any("version", GetVersion)
{{range $CodeDict.TablesInfo}}
    {{.PackageName}}.Init{{.StructName}}RouterGroup(Api){{end}}
}