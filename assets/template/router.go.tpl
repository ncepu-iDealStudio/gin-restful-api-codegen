// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "github.com/gin-gonic/gin"
)

var (
    Api *gin.RouterGroup
)

func Init{{$CodeDict.TableInfo.StructName}}RouterGroup(engine *gin.RouterGroup) {
    Api = engine.Group("{{$CodeDict.TableInfo.PackageName}}")
    Api.GET("", {{$CodeDict.TableInfo.StructName}}GetHandler)
    Api.GET(":{{ if $CodeDict.TableInfo.NaturalKey }}{{ (index $CodeDict.TableInfo.NaturalKey 0) }}{{ else }}{{ (index $CodeDict.TableInfo.Columns 0).Field }}{{ end }}", {{$CodeDict.TableInfo.StructName}}GetBasicHandler)
    Api.POST("", {{$CodeDict.TableInfo.StructName}}PostHandler)
    Api.PUT(":{{ if $CodeDict.TableInfo.NaturalKey }}{{ (index $CodeDict.TableInfo.NaturalKey 0) }}{{ else }}{{ (index $CodeDict.TableInfo.Columns 0).Field }}{{ end }}", {{$CodeDict.TableInfo.StructName}}PutHandler)
    Api.DELETE("", {{$CodeDict.TableInfo.StructName}}DeleteHandler)
    Api.GET("list", {{$CodeDict.TableInfo.StructName}}GetListHandler)
    Api.GET("list/page", {{$CodeDict.TableInfo.StructName}}GetListByPage)
}