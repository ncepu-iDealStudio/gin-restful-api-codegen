// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 18:44
// @Software: GoLand

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs") // 添加搜索路径
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: ", err)
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file:", e.Name, "Op: ", e.Op)
	})
	return
}
