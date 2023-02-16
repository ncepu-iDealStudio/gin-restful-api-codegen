// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:16
// @Software: GoLand

package vipers

import (
	"GinCodeGen/globals/sys"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	genViper *viper.Viper
)

func InitGenViper(confName string) (err error) {
	genViper = viper.New()
	genViper.SetConfigName(confName)
	genViper.AddConfigPath("./configs") // 添加搜索路径
	genViper.SetConfigType("yaml")

	err = genViper.ReadInConfig()
	if err != nil {
		sys.PrintErr("Fatal error config file: ", err)
		return
	}
	genViper.WatchConfig()

	genViper.OnConfigChange(func(e fsnotify.Event) {
		sys.Println("Config file:", e.Name, "Op: ", e.Op)
	})
	return
}

func GetGenViper() *viper.Viper {
	return genViper
}
