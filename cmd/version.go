// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:05
// @Software: GoLand

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",

	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprint(os.Stdout, "GinCodeGen v1.0\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
