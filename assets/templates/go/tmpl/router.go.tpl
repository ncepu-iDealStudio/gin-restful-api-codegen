// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/apis/api_1_0/{{$CodeDict.TableInfo.PackageName}}"
)

var (
    Api *gin.RouterGroup
)

func Init{{$CodeDict.TableInfo.StructName}}RouterGroup(engine *gin.RouterGroup) {
    Api = engine.Group("{{$CodeDict.TableInfo.PackageName}}")
    Api.GET("", {{$CodeDict.TableInfo.PackageName}}.{{$CodeDict.TableInfo.StructName}}GetHandler)
    Api.POST("", {{$CodeDict.TableInfo.PackageName}}.{{$CodeDict.TableInfo.StructName}}PostHandler)
    Api.PUT("", {{$CodeDict.TableInfo.PackageName}}.{{$CodeDict.TableInfo.StructName}}PutHandler)
    Api.DELETE("", {{$CodeDict.TableInfo.PackageName}}.{{$CodeDict.TableInfo.StructName}}DeleteHandler)
    Api.GET("list", {{$CodeDict.TableInfo.PackageName}}.GetListHandler)
    Api.GET("list/page", {{$CodeDict.TableInfo.PackageName}}.GetListByPage)
}