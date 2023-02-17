// coding: utf-8
// @Author : lryself
// @Software: GoLand
{{ $CodeDict := . }}
package dao
import (
    "errors"
    "{{$CodeDict.Dict.ProjectName}}/internal/globals/database"
    "{{$CodeDict.Dict.ProjectName}}/internal/models"
)

type {{$CodeDict.TableInfo.StructName}}Dao struct {
    models.{{$CodeDict.TableInfo.StructName}}Model
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Get(args map[string]any) error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(args).Take(m).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Add(args map[string]any) error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get(args)
    if err == nil {
        return errors.New("数据已存在")
    }
    return mysqlManager.Create(&m).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Update(query map[string]any, args map[string]any) error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get(query)
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Updates(args).Error
}

func (m *{{$CodeDict.TableInfo.StructName}}Dao) Delete(args map[string]any) error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get(args)
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Delete(&m).Error
}