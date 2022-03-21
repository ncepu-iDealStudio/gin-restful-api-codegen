// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:05
// @Software: GoLand

package cmd

import (
	"LRYGoCodeGen/globals/sys"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info.",

	Run: func(cmd *cobra.Command, args []string) {
		sys.Println("v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
