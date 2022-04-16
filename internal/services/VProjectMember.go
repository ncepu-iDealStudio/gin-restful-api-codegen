// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
)

type VProjectMemberService struct {
	dao.VProjectMemberDao
}

func (m *VProjectMemberService) GetList() ([]VProjectMemberService, error) {
	mysqlManager := database.GetMysqlClient()
	results := []VProjectMemberService{}

	rows, err := mysqlManager.Table(m.TableName()).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result VProjectMemberService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *VProjectMemberService) GetListByPage(p parser.ListParser) ([]VProjectMemberService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []VProjectMemberService{}
	for rows.Next() {
		var result VProjectMemberService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
