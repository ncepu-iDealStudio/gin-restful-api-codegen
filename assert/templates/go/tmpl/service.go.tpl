// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package services

import (
    "{{$CodeDict.Dict.ProjectName}}/internal/dao"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/database"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/parser"
)

type {{$CodeDict.TableInfo.StructName}}Service struct {
    dao.{{$CodeDict.TableInfo.StructName}}Dao
}

func (m *{{$CodeDict.TableInfo.StructName}}Service) GetList() ([]{{$CodeDict.TableInfo.StructName}}Service, error) {
    mysqlManager := database.GetMysqlClient()
    results := []{{$CodeDict.TableInfo.StructName}}Service{}

    rows, err := mysqlManager.Table(m.TableName()).Where("IsDeleted", false).Where(m).Rows()
    defer rows.Close()
    if err != nil {
        return results, err
    }
    for rows.Next() {
        var result {{$CodeDict.TableInfo.StructName}}Service
        err = mysqlManager.ScanRows(rows, &result)
        if err != nil {
            return results, err
        }
        results = append(results, result)
    }
    return results, nil
}


func (m *{{$CodeDict.TableInfo.StructName}}Service) GetListByPage(p parser.ListParser) ([]{{$CodeDict.TableInfo.StructName}}Service, error) {
    mysqlManager := database.GetMysqlClient()
    results := []{{$CodeDict.TableInfo.StructName}}Service{}

    rows, err := mysqlManager.Table(m.TableName()).Where("IsDeleted", false).Where(m)
        .Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
    defer rows.Close()
    if err != nil {
        return results, err
    }
    for rows.Next() {
        var result {{$CodeDict.TableInfo.StructName}}Service
        err = mysqlManager.ScanRows(rows, &result)
        if err != nil {
            return results, err
        }
        results = append(results, result)
    }
    return results, nil
}