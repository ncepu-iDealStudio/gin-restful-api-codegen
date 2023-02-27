// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package {{$CodeDict.TableInfo.PackageName}}

import (
    "{{$CodeDict.Dict.ProjectName}}/utils/structs"
    "github.com/gin-gonic/gin"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/parser"
    "{{$CodeDict.Dict.ProjectName}}/internal/services"
    {{if .TableInfo.HasTimeField}} "time" {{end}}
)

{{ if gt $CodeDict.TableInfo.TableType "BASE TABLE" }}func {{$CodeDict.TableInfo.StructName}}GetBasicHandler(c *gin.Context) {
	var err error
    {{ if $CodeDict.TableInfo.NaturalKey }}
	{{ (index $CodeDict.TableInfo.NaturalKey 0) }} := c.Param("{{ (index $CodeDict.TableInfo.NaturalKey 0) }}"{{ else }}{{ (index $CodeDict.TableInfo.Columns 0).Field }} := c.Param("{{ (index $CodeDict.TableInfo.Columns 0).Field }}"{{ end }})

	var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service

	err = {{$CodeDict.TableInfo.StructName}}Service.Get(map[string]any{ {{ if $CodeDict.TableInfo.NaturalKey }}"{{ (index $CodeDict.TableInfo.NaturalKey 0) }}": {{ (index $CodeDict.TableInfo.NaturalKey 0) }}{{ else }}"{{ (index $CodeDict.TableInfo.Columns 0).Field }}": {{ (index $CodeDict.TableInfo.Columns 0).Field }}{{ end }} })

	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}
{{ end }}
func {{$CodeDict.TableInfo.StructName}}GetHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}
        {{.FieldName}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"{{if .NaturalKey}} binding:"required"{{end}}`{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)
    err = {{$CodeDict.TableInfo.StructName}}Service.Get(args)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

{{ if gt $CodeDict.TableInfo.TableType "BASE TABLE" }}func {{$CodeDict.TableInfo.StructName}}PostHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}
        {{.FieldName}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"{{if .NaturalKey}} binding:"required"{{end}}`{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)

    err = {{$CodeDict.TableInfo.StructName}}Service.Add(args)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func {{$CodeDict.TableInfo.StructName}}PutHandler(c *gin.Context) {
    var err error
    {{ if $CodeDict.TableInfo.NaturalKey }}
	{{ (index $CodeDict.TableInfo.NaturalKey 0) }} := c.Param("{{ (index $CodeDict.TableInfo.NaturalKey 0) }}"{{ else }}{{ (index $CodeDict.TableInfo.Columns 0).Field }} := c.Param("{{ (index $CodeDict.TableInfo.Columns 0).Field }}"{{ end }})
    var Parser struct { {{range .TableInfo.Columns}}
        {{.FieldName}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}"`{{end}}
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

	err = {{$CodeDict.TableInfo.StructName}}Service.Update(map[string]any{ {{ if $CodeDict.TableInfo.NaturalKey }}"{{ (index $CodeDict.TableInfo.NaturalKey 0) }}": {{ (index $CodeDict.TableInfo.NaturalKey 0) }}{{ else }}"{{ (index $CodeDict.TableInfo.Columns 0).Field }}": {{ (index $CodeDict.TableInfo.Columns 0).Field }}{{ end }} }, args)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

func {{$CodeDict.TableInfo.StructName}}DeleteHandler(c *gin.Context) {
    var err error

    var Parser struct { {{range .TableInfo.Columns}}{{if .NaturalKey}}
        {{.FieldName}} {{.DataType}} `json:"{{.Field}}" form:"{{.Field}}" binding:"required"`{{end}}{{end}}
    }
    err = c.ShouldBind(&Parser)
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    var {{$CodeDict.TableInfo.StructName}}Service services.{{$CodeDict.TableInfo.StructName}}Service
    {{$CodeDict.TableInfo.StructName}}Service.Assign(Parser)

    args, err := structs.StructToMap(Parser, "json")
    if err != nil {
        parser.JsonParameterIllegal(c, "", err)
        return
    }

    err = {{$CodeDict.TableInfo.StructName}}Service.Delete(args)
    if err != nil {
        parser.JsonDBError(c, "", err)
        return
    }
    parser.JsonOK(c, "", {{$CodeDict.TableInfo.StructName}}Service)
}

{{ end }}

// 获取列表
func {{$CodeDict.TableInfo.StructName}}GetListHandler(c *gin.Context) {
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
func {{$CodeDict.TableInfo.StructName}}GetListByPage(c *gin.Context) {
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
        "totalPage":  int(count) / Parser.Size + 1,
    })
}