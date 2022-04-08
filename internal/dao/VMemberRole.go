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
    return mysqlManager.Where("IsDeleted", false).Where(m).Take(m).Error
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
    return mysqlManager.Model(&m).Update("IsDeleted", true).Error
}