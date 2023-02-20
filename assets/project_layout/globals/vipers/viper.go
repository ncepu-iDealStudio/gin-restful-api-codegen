// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:16
// @Software: GoLand

package vipers

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"tem_go_project/globals/sys"
)

var (
	tempViper *viper.Viper
)

func InitTempViper(confName string) (err error) {
	tempViper = viper.New()
	tempViper.SetConfigName(confName)
	tempViper.AddConfigPath("./configs") // 添加搜索路径
	tempViper.SetConfigType("yaml")

	err = tempViper.ReadInConfig()
	if err != nil {
		sys.PrintErr("Fatal error config file: ", err)
		return
	}
	tempViper.WatchConfig()

	tempViper.OnConfigChange(func(e fsnotify.Event) {
		sys.Println("Config file:", e.Name, "Op: ", e.Op)
	})
	return
}

func GetTempViper() *viper.Viper {
	return tempViper
}
