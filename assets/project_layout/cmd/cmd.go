// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 22:40
// @Software: GoLand

package cmd

import (
	"github.com/spf13/cobra"
	"tem_go_project/apis"
	"tem_go_project/utils"
	"tem_go_project/utils/errHelper"
	"tem_go_project/utils/message"
)

var rootCmd = &cobra.Command{
	Use:   "project_layout",
	Short: "tem_go_project服务",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Start doing things.开始做事情
		message.Println("开始启动服务！")
		utils.GetWaitGroup().Add(1)
		go apis.StartHttp()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	errHelper.ErrExit(rootCmd.Execute())
}

func init() {
	//rootCmd.PersistentFlags().StringP("Port", "P", "8000", "配置文件名(注意-C为大写)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	//port, err := rootCmd.Flags().GetString("Port")
	//errHelper.ErrExit(err)
}
