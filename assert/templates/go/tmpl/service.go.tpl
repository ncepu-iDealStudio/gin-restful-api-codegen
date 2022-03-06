// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package services

import (
    "{{$CodeDict.Dict.ProjectName}}/internal/dao"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/database"
)

type {{$CodeDict.TableInfo.StructName}}Service struct {
    dao.{{$CodeDict.TableInfo.StructName}}Dao
}

func (m *{{$CodeDict.TableInfo.StructName}}Service) GetList() ([]{{$CodeDict.TableInfo.StructName}}Service, error) {
    mysqlManager := database.GetMysqlClient()

    rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
        "IsDeleted": 0,
    }).Where(m).Rows()
    defer rows.Close()
    if err != nil {
        return nil, err
    }
    results := []{{$CodeDict.TableInfo.StructName}}Service{}
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