// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 20:40
// @Software: GoLand

package settings

import (
	logs "gitee.com/lryself/go-utils/loggers"
	"github.com/spf13/viper"
	"LRYGoCodeGen/internal/globals/database"
)

func InitDatabase() (err error) {
	var log = logs.GetLogger()
	if viper.GetBool("system.UseMysql") {
		err = database.InitMysqlClient()
		if err != nil {
			log.Errorln("mysql初始化出错:", err)
			return
		}
	}
	if viper.GetBool("system.UseRedis") {
		err = database.InitRedisClient()
		if err != nil {
			log.Errorln("redis初始化出错:", err)
			return
		}
	}
	return
}
