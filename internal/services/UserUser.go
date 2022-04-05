// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
)

type UserUserService struct {
	dao.UserUserDao
}

func (m *UserUserService) GetList() ([]UserUserService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []UserUserService{}
	for rows.Next() {
		var result UserUserService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *UserUserService) GetListByPage(p parser.ListParser) ([]UserUserService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []UserUserService{}
	for rows.Next() {
		var result UserUserService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
