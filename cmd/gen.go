// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 22:59
// @Software: GoLand

package cmd

import (
	"GinCodeGen/core/gen/gen_db"
	"GinCodeGen/core/gen/gen_program"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
		_, _ = fmt.Fprint(os.Stdout, "开始生成代码！\n")
		for _, arg := range args {
			if arg == "p" {
				gen_program.GenProgramCodeFromTemplates()
			} else if arg == "d" {
				gen_db.GenDBCodeFromTemplate()
			}
		}
		_, _ = fmt.Fprint(os.Stdout, "生成结束！\n")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
