// coding: utf-8
// @Author : lryself
// @Date : 2022/3/31 21:29
// @Software: GoLand

package cmd

import (
	"LRYGoCodeGen/internal"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "开启http服务",

	Run: func(cmd *cobra.Command, args []string) {
		internal.StartHttp()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
