// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
	"LRYGoCodeGen/internal/dao"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/parser"
	"gorm.io/gorm"
)

type ProjectPoolService struct {
	dao.ProjectPoolDao
}

func (m *ProjectPoolService) GetList() ([]ProjectPoolService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []ProjectPoolService{}
	for rows.Next() {
		var result ProjectPoolService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *ProjectPoolService) GetListByPage(p parser.ListParser) ([]ProjectPoolService, error) {
	mysqlManager := database.GetMysqlClient()

	rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Limit(p.Limit).Offset(p.Offset).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := []ProjectPoolService{}
	for rows.Next() {
		var result ProjectPoolService
		err = mysqlManager.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (m *ProjectPoolService) Add() error {
	var err error
	mysqlManager := database.GetMysqlClient()
	err = mysqlManager.Transaction(func(tx *gorm.DB) error {
		err = mysqlManager.Create(&m).Error
		if err != nil {
			return err
		}
		var projectMember ProjectMemberService
		projectMember.ProjectID = m.ProjectID
		projectMember.UserID = m.UserID
		projectMember.RoleID = "1"
		err = mysqlManager.Create(&projectMember).Error
		if err != nil {
			return err
		}
		return err
	})
	return err
}
