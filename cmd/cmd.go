// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 22:40
// @Software: GoLand

package cmd

import (
	"LRYGoCodeGen/core"
	"LRYGoCodeGen/globals/sys"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lryCG",
	Short: "gen code with templates tools",
	Long:  `base on templates for project or database to code`,
	Run: func(cmd *cobra.Command, args []string) {
		// Start doing things.开始做事情
		sys.Println("开始执行默认方法！")
		core.Execute()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
