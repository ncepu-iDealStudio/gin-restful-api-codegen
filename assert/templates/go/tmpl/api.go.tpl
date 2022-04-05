// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/codes"
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

    var {{$CodeDict.TableInfo.StructName}} services.{{$CodeDict.TableInfo.StructName}}Service
    //针对业务主键处理{{range $CodeDict.TableInfo.NaturalKey}}
    {{$CodeDict.TableInfo.StructName}}.{{.}} = {{$CodeDict.TableInfo.StructName}}Service.{{.}}{{end}}

    if c.Request.Method == "GET" {
        err = {{$CodeDict.TableInfo.StructName}}Service.Get()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
    } else if c.Request.Method == "POST" {
        err = {{$CodeDict.TableInfo.StructName}}.Get()
        if err != nil {
            if err.Error() != "record not found" {
                parser.JsonDBError(c, "", err)
                return
            }
        } else {
            c.JSON(http.StatusOK, gin.H{
                "code":    codes.DataExist,
                "message": "数据已存在！",
            })
            return
        }

        err = {{$CodeDict.TableInfo.StructName}}Service.Add()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
    } else if c.Request.Method == "PUT" {
        args, err := {{$CodeDict.TableInfo.StructName}}Service.GetModelMap()
        if err != nil {
            parser.JsonParameterIllegal(c, "", err)
            return
        }
        //不能修改业务主键{{range $CodeDict.TableInfo.NaturalKey}}
        delete(args, "{{.}}"){{end}}

        err = {{$CodeDict.TableInfo.StructName}}.Update(args)
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
        {{$CodeDict.TableInfo.StructName}}Service = {{$CodeDict.TableInfo.StructName}}
    } else if c.Request.Method == "DELETE" {
        err = {{$CodeDict.TableInfo.StructName}}.Delete()
        if err != nil {
            parser.JsonDBError(c, "", err)
            return
        }
        {{$CodeDict.TableInfo.StructName}}Service.IsDeleted = true
    }
        
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}
// 获取列表
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

// 获取列表（分页）
func GetListByPage(c *gin.Context) {
    var err error

    var Parser struct {
        services.{{$CodeDict.TableInfo.StructName}}Service
        parser.ListParser
    }

    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    results, err := Parser.GetListByPage(Parser.ListParser)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }

    parser.JsonOK(c, "", results)
}