// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "gitee.com/lryself/go-utils/structs"
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/parser"
    "{{$CodeDict.Dict.ProjectName}}/internal/services"
    "time"
)

func {{$CodeDict.TableInfo.StructName}}PostHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}
        {{.Field}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"{{if .NaturalKey}} binding:"required"{{end}}`{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)

    err = {{$CodeDict.TableInfo.StructName}}Service.Add()
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func {{$CodeDict.TableInfo.StructName}}GetHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}
        {{.Field}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"{{if .NaturalKey}} binding:"required"{{end}}`{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)
    err = {{$CodeDict.TableInfo.StructName}}Service.Get()
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func {{$CodeDict.TableInfo.StructName}}PutHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}
        {{.Field}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"{{if .NaturalKey}} binding:"required"{{end}}`{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service

    args, err := structs.StructToMap(Parser, "json")
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }
    //不能修改业务主键{{range $CodeDict.TableInfo.NaturalKey}}
    delete(args, "{{.}}"){{end}}

    err = {{$CodeDict.TableInfo.StructName}}Service.Update(args)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func {{$CodeDict.TableInfo.StructName}}DeleteHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}{{if .NaturalKey}}
        {{.Field}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}" binding:"required"`{{end}}{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)

    err = {{$CodeDict.TableInfo.StructName}}Service.Delete()
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
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

    results, count, err := Parser.GetListByPage(Parser.ListParser)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }

    parser.JsonOK(c, "", results, map[string]any{
        "totalCount": count,
        "totalPage":  int(count) / Parser.Size,
    })
}