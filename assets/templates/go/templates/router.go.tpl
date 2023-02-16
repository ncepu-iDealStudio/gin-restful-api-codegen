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
    Api.POST("", {{$CodeDict.TableInfo.StructName}}PostHandler)
    Api.PUT("", {{$CodeDict.TableInfo.StructName}}PutHandler)
    Api.DELETE("", {{$CodeDict.TableInfo.StructName}}DeleteHandler)
    Api.GET("list", {{$CodeDict.TableInfo.StructName}}GetListHandler)
    Api.GET("list/page", {{$CodeDict.TableInfo.StructName}}GetListByPage)
}