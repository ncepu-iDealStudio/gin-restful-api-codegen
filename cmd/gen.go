// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 22:59
// @Software: GoLand

package cmd

import (
	"LRYGoCodeGen/core/gen/gen_db"
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/globals/sys"
	"errors"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成代码命令，可选(p,d)",
	Long:  "生成代码命令，可选(p,d)\np —— 生成项目代码\nd —— 生成数据库代码",
	Args: func(cmd *cobra.Command, args []string) error {
		for _, v := range args {
			if !func(s string) bool {
				for _, arg := range []string{"p", "d"} {
					if s == arg {
						return true
					}
				}
				return false
			}(v) {
				return errors.New("输入范围应在(p,d)之中！")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sys.Println("开始生成代码！")
		for _, arg := range args {
			if arg == "p" {
				gen_program.GenProgramCodeFromTemplates()
			} else if arg == "d" {
				gen_db.GenDBCodeFromTemplate()
			}
		}
		sys.Println("生成结束！")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
