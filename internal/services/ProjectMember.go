// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
    "LRYGoCodeGen/internal/dao"
    "LRYGoCodeGen/internal/globals/database"
)

type ProjectMemberService struct {
    dao.ProjectMemberDao
}

func (m *ProjectMemberService) GetList() ([]ProjectMemberService, error) {
    mysqlManager := database.GetMysqlClient()

    rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
        "IsDeleted": 0,
    }).Where(m).Rows()
    defer rows.Close()
    if err != nil {
        return nil, err
    }
    results := []ProjectMemberService{}
    for rows.Next() {
        var result ProjectMemberService
        err = mysqlManager.ScanRows(rows, &result)
        if err != nil {
            return results, err
        }
        results = append(results, result)
    }
    return results, nil
}