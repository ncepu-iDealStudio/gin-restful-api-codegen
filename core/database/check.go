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
		err = errors.New("be care! database name is a keyword! ")
		return
	}
	return nil
}

func CheckTable(table mysql.TableModel) (err error) {
	tableName := table.TableName
	// 判断表名是否是关键字
	if keywords.In(tableName) {
		msg := fmt.Sprintf("be care!  table name '%s' is a keyword, you should change it! We skip this table during generating code.", tableName)
		err = errors.New(msg)
		return
	}
	columns := table.Columns
	for _, colum := range columns {
		// 判断列名是否是关键字
		if keywords.In(colum.Field) {
			msg := fmt.Sprintf(
				"be care!  column name '%s.%s' is a keyword, you should change it! We skip this table during generating code.",
				tableName, colum.Field,
			)
			err = errors.New(msg)
			return
		}
	}
	return
}

// CheckPrimaryKey 检查目标表是否有主键
func CheckPrimaryKey(table mysql.TableModel) (err error) {
	var havePrimaryKey = false
	if table.TableType == "VIEW" {
		return nil
	}
	for _, column := range table.Columns {
		if column.Key == "PRI" {
			havePrimaryKey = true
			break
		}
		if !havePrimaryKey {
			msg := fmt.Sprintf(
				"be care! Table '%s' does not contain a primary key. We skip this table during generating code.",
				table.TableName,
			)
			err = errors.New(msg)
			return
		}
	}
	return nil
}
