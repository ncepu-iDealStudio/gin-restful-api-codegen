// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 15:28
// @Software: GoLand

package mysql

import (
	"GinCodeGen/globals/vipers"
	"fmt"
	"sort"
	"strings"
)

type DataBaseModel struct {
	DataBaseName string
	Tables       []TableModel
}

type TableModel struct {
	TableName string
	Comment   string
	Columns   []columnModel
}

type columnModel struct {
	Field      string
	Type       string
	Collation  *string
	Null       bool
	Key        string
	Default    *string
	Extra      string
	Privileges []string
	Comment    string
}

func (m *DataBaseModel) GetTables(orm *MySQLOrm) error {
	var err error
	tables, err := orm.GetTables(m.DataBaseName)
	if err != nil {
		return err
	}
	for k, v := range tables {
		var table TableModel
		table.TableName = k
		table.Comment = v
		table.GetColumns(orm)
		m.Tables = append(m.Tables, table)
	}
	sort.Slice(m.Tables, func(i, j int) bool {
		return m.Tables[i].TableName < m.Tables[j].TableName
	})
	return nil
}

func (m *TableModel) GetColumns(orm *MySQLOrm) {
	columns, err := orm.GetColumns(m.TableName)
	if err != nil {
		return
	}
	for _, column := range columns {
		c := columnModel{
			Field:     column.Field,
			Type:      column.Type,
			Collation: column.Collation,
			Key:       column.Key,
			Default:   column.Default,
			Extra:     column.Extra,
			Comment:   column.Comment,
		}
		if column.Null == "YES" {
			c.Null = true
		} else {
			c.Null = false
		}
		if column.Privileges != "" {
			c.Privileges = strings.Split(column.Privileges, ",")
		}
		m.Columns = append(m.Columns, c)
	}
	return
}

func GetMysqlDBModel() (*DataBaseModel, error) {
	var genViper = vipers.GetGenViper()
	var DBModel DataBaseModel
	var err error
	orm, err := InitDBOrm(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		genViper.GetString("database.username"),
		genViper.GetString("database.password"),
		genViper.GetString("database.host"),
		genViper.GetString("database.port"),
		genViper.GetString("database.database"),
	))
	if err != nil {
		return nil, err
	}
	defer orm.DestroyDB()
	DBModel.DataBaseName = genViper.GetString("database.database")
	err = DBModel.GetTables(orm)
	if err != nil {
		return nil, err
	}
	return &DBModel, nil
}
