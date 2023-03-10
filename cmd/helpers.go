// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:04
// @Software: GoLand

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func Error(cmd *cobra.Command, args []string, err error) {
	_, _ = fmt.Fprint(os.Stderr, fmt.Sprintf("execute %s args:%v error:%v\n", cmd.Name(), args, err))
	os.Exit(1)
}

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}
