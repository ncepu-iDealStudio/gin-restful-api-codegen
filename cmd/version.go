// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:05
// @Software: GoLand

package cmd

import (
	"GinCodeGen/globals/sys"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",

	Run: func(cmd *cobra.Command, args []string) {
		sys.Println("v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
