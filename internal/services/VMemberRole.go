// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
)

type VMemberRoleService struct {
	dao.VMemberRoleDao
}

func (m *VMemberRoleService) GetList() ([]VMemberRoleService, error) {
	mysqlManager := database.GetMysqlClient()
	results := []VMemberRoleService{}

	rows, err := mysqlManager.Table(m.TableName()).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result VMemberRoleService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *VMemberRoleService) GetListByPage(p parser.ListParser) ([]VMemberRoleService, error) {
	mysqlManager := database.GetMysqlClient()
	results := []VMemberRoleService{}

	rows, err := mysqlManager.Table(m.TableName()).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result VMemberRoleService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
