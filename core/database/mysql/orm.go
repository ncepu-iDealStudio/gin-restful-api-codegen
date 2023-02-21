// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 16:01
// @Software: GoLand

package mysql

import (
	"GinCodeGen/tools/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type MySQLOrm struct {
	DB *sql.DB
}

// InitDBOrm init MySqlDB
func InitDBOrm(dataSource string) (orm *MySQLOrm, err error) {
	orm = new(MySQLOrm)
	err = orm.GetDBOrm(dataSource)
	return
}

// GetDBOrm get gorm.db
func (orm *MySQLOrm) GetDBOrm(dataSourceName string) (err error) {
	if orm.DB == nil {
		orm.DB, err = sql.Open("mysql", dataSourceName)
		//orm.DB, err = gorm.Open(
		//	mysql.Open(dataSourceName),
		//	&gorm.Config{
		//		PrepareStmt:    false,
		//		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 全局禁用表名复数
		//		Logger:         logger.Default,
		//	})
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, "please check your database config！\n")
			log := logger.GetLogger()
			log.Error("please check your database config！\n")
			return
		}
	}
	return
}

// DestroyDB destroydb
func (orm *MySQLOrm) DestroyDB() {
	if orm.DB != nil {
		orm.DB.Close()
		orm.DB = nil
	}
}

func (orm MySQLOrm) Raw(query string) (*sql.Rows, error) {
	rows, err := orm.DB.Query(query)
	return rows, err
}

func (orm *MySQLOrm) GetTablesName() ([]string, error) {
	var tables []string
	var err error
	rows, err := orm.DB.Query("show tables")
	defer rows.Close()
	if err != nil {
		log := logger.GetLogger()
		log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			log := logger.GetLogger()
			log.Error(err)
			log.Error(fmt.Sprintf("error occur during scaning rows of a table: %s", err))
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (orm *MySQLOrm) GetTables(dbName string) (map[string]string, error) {
	result := map[string]string{}
	var err error
	rows, err := orm.Raw("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.TABLES WHERE table_schema= '" + dbName + "'")
	defer func(rows *sql.Rows) {
		if rows != nil {
			err = rows.Close()
			if err != nil {
				_, _ = fmt.Fprint(os.Stderr, fmt.Sprintln(err))
				log := logger.GetLogger()
				log.Error(fmt.Sprintf("unable to get tables: %s", err))
			}
		}
	}(rows)
	if err != nil {
		log := logger.GetLogger()
		log.Error(err)
		return map[string]string{}, err
	}

	for rows.Next() {
		var table, comment string
		err1 := rows.Scan(&table, &comment)
		if err1 != nil {
			log := logger.GetLogger()
			log.Error(fmt.Sprintf("error occur during scaning rows of a table: %s", err))
			return nil, err1
		}
		result[table] = comment
	}
	return result, nil
}

// genColumns show full columns
type columns struct {
	Field      string
	Type       string
	Collation  *string
	Null       string
	Key        string
	Default    *string
	Extra      string
	Privileges string
	Comment    string
}

func (orm *MySQLOrm) GetColumns(tableName string) ([]columns, error) {
	var result []columns
	rows, err := orm.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c columns
		err1 := rows.Scan(&c.Field, &c.Type, &c.Collation, &c.Null, &c.Key, &c.Default, &c.Extra, &c.Privileges, &c.Comment)
		if err1 != nil {
			log := logger.GetLogger()
			log.Error(err)
			return nil, err1
		}
		result = append(result, c)
	}
	return result, err
}
