// coding: utf-8
// @Author : lryself
// @Software: GoLand

package dao
import (
    "errors"
    "LRYGoCodeGen/internal/globals/database"
    "LRYGoCodeGen/internal/models/mysqlModel"
)

type VMemberRoleDao struct {
    mysqlModel.VMemberRoleModel
}

func (m *VMemberRoleDao) Get() error {
    mysqlManager := database.GetMysqlClient()
    return mysqlManager.Where(map[string]interface{}{
        "IsDeleted": 0,
    }).Where(m).Take(m).Error
}

func (m *VMemberRoleDao) Add() error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err == nil {
        return errors.New("数据已存在")
    }
    return mysqlManager.Create(&m).Error
}

func (m *VMemberRoleDao) Update(args map[string]interface{}) error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Updates(args).Error
}

func (m *VMemberRoleDao) Delete() error {
    mysqlManager := database.GetMysqlClient()
    err := m.Get()
    if err != nil {
        return err
    }
    return mysqlManager.Model(&m).Updates(map[string]interface{}{
        "IsDeleted": 1,
    }).Error
}