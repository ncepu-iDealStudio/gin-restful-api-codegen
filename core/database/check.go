package database

import (
	"GinCodeGen/core/database/mysql"
	"GinCodeGen/tools/common/keywords"
	"errors"
	"fmt"
)

// CheckDatabase 关键字检查-数据库名称，表名称，字段名称不能为Go的关键字
func CheckDatabase(model *mysql.DataBaseModel) (err error) {
	// 判断库名是否为关键字
	dbName := model.DataBaseName
	if keywords.In(dbName) {
		err = errors.New("database name is keyword")
		return
	}

	tables := model.Tables
	for _, elem := range tables {
		tableName := elem.TableName
		// 判断表名是否是关键字
		if keywords.In(tableName) {
			msg := fmt.Sprintf("table name '%s' is keyword, you should change it.", tableName)
			err = errors.New(msg)
			return
		}
		columns := elem.Columns
		for _, colum := range columns {
			// 判断列名是否是关键字
			if keywords.In(colum.Field) {
				msg := fmt.Sprintf(
					"column name '%s.%s' is keyword, you should change it.",
					tableName, colum.Field,
				)
				err = errors.New(msg)
				return
			}
		}
	}
	return nil
}

// CheckPrimaryKey 检查目标表是否有主键
func CheckPrimaryKey(model *mysql.DataBaseModel) (err error) {
	var havePrimaryKey bool
	tables := model.Tables
	for _, table := range tables {
		havePrimaryKey = false
		if table.TableType == "VIEW" {
			continue
		}
		for _, column := range table.Columns {
			if column.Key == "PRI" {
				havePrimaryKey = true
				break
			}
		}

		if !havePrimaryKey {
			msg := fmt.Sprintf(
				"Table '%s' does not contain a primary key.",
				table.TableName,
			)
			err = errors.New(msg)
			return
		}
	}
	return nil
}
