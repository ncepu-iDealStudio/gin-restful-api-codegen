// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 15:58
// @Software: GoLand

package genmysql

import "LRYGoCodeGen/core/globals/config"

func GetMysqlDBModel() (*DataBaseModel, error) {
	var DBModel DataBaseModel
	var err error
	dbConfig, err := config.GetDBConfig()
	if err != nil {
		return nil, err
	}
	orm, err := InitDBOrm(dbConfig.GetMysqlConnConfig())
	if err != nil {
		return nil, err
	}
	defer orm.DestroyDB()
	DBModel.DataBaseName = dbConfig.Mysql.Database
	err = DBModel.GetTables(orm)
	if err != nil {
		return nil, err
	}
	return &DBModel, nil
}
