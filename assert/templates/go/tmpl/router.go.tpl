// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/apis/api1_0/{{$CodeDict.TableInfo.PackageName}}"
)

var (
    Api *gin.RouterGroup
)

func Init{{$CodeDict.TableInfo.StructName}}RouterGroup(engine *gin.RouterGroup) {
    Api = engine.Group("{{$CodeDict.TableInfo.PackageName}}")
    Api.Any("", {{$CodeDict.TableInfo.PackageName}}.{{$CodeDict.TableInfo.StructName}}Api)
    Api.GET("list", {{$CodeDict.TableInfo.PackageName}}.GetListHandler)
}