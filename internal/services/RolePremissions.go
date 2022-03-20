// coding: utf-8
// @Author : lryself
// @Software: GoLand

package services

import (
    "LRYGoCodeGen/internal/dao"
    "LRYGoCodeGen/internal/globals/database"
)

type RolePremissionsService struct {
    dao.RolePremissionsDao
}

func (m *RolePremissionsService) GetList() ([]RolePremissionsService, error) {
    mysqlManager := database.GetMysqlClient()

    rows, err := mysqlManager.Table(m.TableName()).Where(map[string]interface{}{
        "IsDeleted": 0,
    }).Where(m).Rows()
    defer rows.Close()
    if err != nil {
        return nil, err
    }
    results := []RolePremissionsService{}
    for rows.Next() {
        var result RolePremissionsService
        err = mysqlManager.ScanRows(rows, &result)
        if err != nil {
            return results, err
        }
        results = append(results, result)
    }
    return results, nil
}