package init

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
	"tem_go_project/utils"
	"tem_go_project/utils/loggers"
	"tem_go_project/utils/message"
)

func Init() *sync.WaitGroup {
	// 消息初始化
	waitGroup := MessageInit()

	// 配置读取初始化
	err := ViperInit()
	if err != nil {
		message.PrintErr(err)
		message.Exit()
	}
	message.Println("配置文件加载完成")

	// 日志初始化
	LoggerInit(viper.GetString("log.type"))
	message.Println("日志组件初始化完成")

	return waitGroup
}

func MessageInit() *sync.WaitGroup {
	waitGroup := utils.GetWaitGroup()
	waitGroup.Add(1)
	go message.InitMsg()
	return waitGroup
}

func ViperInit() (err error) {
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

func LoggerInit(logType string) {
	loggers.InitLogger(logType)
}
