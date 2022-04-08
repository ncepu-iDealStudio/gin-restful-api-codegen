// coding: utf-8
// @Author : lryself
// @Software: GoLand

package dao

import (
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/models/mysqlModel"
	"errors"
)

type SsoUserDao struct {
	mysqlModel.SsoUserModel
}

func (m *SsoUserDao) Get() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where("IsDeleted", false).Where(m).Take(m).Error
}

func (m *SsoUserDao) Add() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err == nil {
		return errors.New("数据已存在")
	}
	return mysqlManager.Create(&m).Error
}

func (m *SsoUserDao) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *SsoUserDao) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Update("IsDeleted", true).Error
}
