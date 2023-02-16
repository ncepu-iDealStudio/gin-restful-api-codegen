// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 22:40
// @Software: GoLand

package cmd

import (
	"GinCodeGen/core"
	"GinCodeGen/globals/sys"
	"GinCodeGen/globals/vipers"
	"GinCodeGen/utils/errHelper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lryCG",
	Short: "生成代码的工具",
	Long:  `基于模板生成项目代码的工具`,
	Run: func(cmd *cobra.Command, args []string) {
		// Start doing things.开始做事情
		sys.Println("开始生成代码！")
		core.Execute()
		sys.Println("代码生成完成！")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	errHelper.ErrExit(rootCmd.Execute())
}
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("Config", "C", "config", "配置文件名(注意-C为大写)")
}

func initConfig() {
	confName, err := rootCmd.Flags().GetString("Config")
	errHelper.ErrExit(err)
	errHelper.ErrExit(vipers.InitGenViper(confName))
}
