// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
)

type RolePermissionsService struct {
	dao.RolePermissionsDao
}

func (m *RolePermissionsService) GetList() ([]RolePermissionsService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []RolePermissionsService{}
	for rows.Next() {
		var result RolePermissionsService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *RolePermissionsService) GetListByPage(p parser.ListParser) ([]RolePermissionsService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []RolePermissionsService{}
	for rows.Next() {
		var result RolePermissionsService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
