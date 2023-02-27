// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 20:24
// @Software: GoLand

package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tem_go_project/internal/globals/vipers"
)

var (
	mysqlClient *gorm.DB
)

func GetMysqlClient() *gorm.DB {
	if gin.Mode() == gin.ReleaseMode {
		return mysqlClient
	}
	return mysqlClient.Debug()
}

func InitMysqlClient() (err error) {
	v := vipers.GetDatabaseViper()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		v.GetString("mysql.username"),
		v.GetString("mysql.password"),
		v.GetString("mysql.host"),
		v.GetString("mysql.port"),
		v.GetString("mysql.database"),
	)
	mysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
