// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:16
// @Software: GoLand

package vipers

import (
	"LRYGoCodeGen/globals/sys"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(confName string) (err error) {
	viper.SetConfigName(confName)
	viper.AddConfigPath("./configs") // 添加搜索路径
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		sys.PrintErr("Fatal error config file: ", err)
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		sys.Println("Config file:", e.Name, "Op: ", e.Op)
	})
	return
}
