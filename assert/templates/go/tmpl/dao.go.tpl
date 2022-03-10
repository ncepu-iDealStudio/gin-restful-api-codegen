// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package dao
import (
    "errors"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/database"
    "{{$CodeDict.Dict.ProjectName}}/internal/models/mysqlModel"
)

type {{$CodeDict.TableInfo.StructName}}Dao struct {
    mysqlModel.{{$CodeDict.TableInfo.StructName}}Model
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Get() error {
    mysqlManager := database.GetMysqlClient()
    return mysqlManager.Where(map[string]interface{}{
        "IsDeleted": 0,
    }).Where(m).Take(m).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Add() error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err == nil {
        return errors.New("数据已存在")
    }
    return mysqlManager.Create(&m).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Update(args map[string]interface{}) error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Updates(args).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Delete() error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Updates(map[string]interface{}{
        "IsDeleted": 1,
    }).Error
}