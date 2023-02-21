// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 22:40
// @Software: GoLand

package cmd

import (
	"GinCodeGen/core"
	initialization "GinCodeGen/init"
	"GinCodeGen/tools/errorPack"
	"GinCodeGen/tools/logger"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lryCG",
	Short: "生成代码的工具",
	Long:  `基于模板生成项目代码的工具`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.GetLogger()
		// Start doing things.开始做事情
		log.Infoln("代码生成开始执行")
		_, _ = fmt.Fprint(os.Stdout, "开始生成代码！\n")
		core.Execute()
		_, _ = fmt.Fprint(os.Stdout, "代码生成完成！\n")
		log.Infoln("代码生成完成")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	errorPack.ErrExit(rootCmd.Execute())
}
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("Config", "C", "config", "配置文件名(注意-C为大写)")
}

func initConfig() {
	var log = logger.GetLogger()
	confName, err := rootCmd.Flags().GetString("Config")
	errorPack.ErrExit(err)
	log.Infoln(fmt.Sprintf("配置文件初始化，文件地址：config/%s.yaml", confName))
	errorPack.ErrExit(initialization.InitCodeGenViper(confName))
	log.Infoln("配置文件初始化完成！")
}
