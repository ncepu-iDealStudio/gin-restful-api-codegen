// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
)

type SsoUserService struct {
	dao.SsoUserDao
}

func (m *SsoUserService) GetList() ([]SsoUserService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []SsoUserService{}
	for rows.Next() {
		var result SsoUserService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *SsoUserService) GetListByPage(p parser.ListParser) ([]SsoUserService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []SsoUserService{}
	for rows.Next() {
		var result SsoUserService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}