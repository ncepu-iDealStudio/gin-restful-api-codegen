// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 16:06
// @Software: GoLand

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBConfig struct {
	Mysql dbInfo `yaml:"mysql"`
}

type dbInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func GetDBConfig() (*DBConfig, error) {
	var dbConfig DBConfig
	var err error
	file, err := ioutil.ReadFile("configs/database.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &dbConfig)
	if err != nil {
		return nil, err
	}
	return &dbConfig, nil
}

func (d DBConfig) GetMysqlConnConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		d.Mysql.Username,
		d.Mysql.Password,
		d.Mysql.Host,
		d.Mysql.Port,
		d.Mysql.Database,
	)
}
