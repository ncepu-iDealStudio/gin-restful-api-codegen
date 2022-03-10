// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 15:28
// @Software: GoLand

package gen_mysql

import (
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
