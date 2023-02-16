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

type CodeGenVipers struct {
	genViper *viper.Viper
	sysViper *viper.Viper
}

var CodeGenViper CodeGenVipers

func InitViper(confName string) (_viper *viper.Viper, err error) {
	_viper = viper.New()
	_viper.SetConfigName(confName)
	_viper.AddConfigPath("./configs") // 添加搜索路径
	_viper.SetConfigType("yaml")

	err = _viper.ReadInConfig()
	if err != nil {
		sys.PrintErr("Fatal error config file: ", err)
		return nil, err
	}
	_viper.WatchConfig()

	_viper.OnConfigChange(func(e fsnotify.Event) {
		sys.Println("Config file:", e.Name, "Op: ", e.Op)
	})
	return _viper, nil
}

func (this *CodeGenVipers) InitGenViper(confName string) (err error) {
	this.genViper, err = InitViper(confName)
	if err != nil {
		return
	}
	return
}

func (this *CodeGenVipers) InitSysViper(confName string) (err error) {
	this.sysViper, err = InitViper(confName)
	if err != nil {
		return
	}
	return
}

func (this *CodeGenVipers) GetGenViper() *viper.Viper {
	return this.genViper
}

func (this *CodeGenVipers) GetSysViper() *viper.Viper {
	return this.sysViper
}

func InitCodeGenViper(genConfigName string) (err error) {
	CodeGenViper = CodeGenVipers{}
	err = CodeGenViper.InitGenViper(genConfigName)
	if err != nil {
		return
	}
	err = CodeGenViper.InitSysViper("sys")
	if err != nil {
		return
	}
	return
}

func GetCodeGenViper() *CodeGenVipers {
	return &CodeGenViper
}
