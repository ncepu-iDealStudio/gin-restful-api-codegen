// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "github.com/gin-gonic/gin"
    "gitee.com/lryself/go-utils/structs"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/parser"
    "{{$CodeDict.Dict.ProjectName}}/internal/services"
)

func {{$CodeDict.TableInfo.StructName}}Api(c *gin.Context) {
    var err error
    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service

    err = c.ShouldBind(&{{$CodeDict.TableInfo.StructName}}Service)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }
    
    if c.Request.Method == "GET" {
        err = {{$CodeDict.TableInfo.StructName}}Service.Get()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
    } else if c.Request.Method == "POST" {
        err = {{$CodeDict.TableInfo.StructName}}Service.Add()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
    } else if c.Request.Method == "PUT" {
        args, err := structs.StructToMap({{$CodeDict.TableInfo.StructName}}Service.{{$CodeDict.TableInfo.StructName}}Dao.{{$CodeDict.TableInfo.StructName}}Model, "json")
        if err != nil {
            parser.JsonParameterIllegal(c, "", err)
            return
        }
        // todo 请将AutoID修改为业务主键名
        delete(args, "AutoID")
        temp := services.{{$CodeDict.TableInfo.StructName}}Service{}
        temp.AutoID = {{$CodeDict.TableInfo.StructName}}Service.AutoID

        err = temp.Update(args)
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
    } else if c.Request.Method == "DELETE" {
        err = {{$CodeDict.TableInfo.StructName}}Service.Delete()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
        {{$CodeDict.TableInfo.StructName}}Service.IsDeleted = true
    }
        
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func GetListHandler(c *gin.Context) {
    var err error
    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service

    err = c.ShouldBind(&{{$CodeDict.TableInfo.StructName}}Service)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }
    
    results, err := {{$CodeDict.TableInfo.StructName}}Service.GetList()
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    
    parser.JsonOK(c, "", results)
}