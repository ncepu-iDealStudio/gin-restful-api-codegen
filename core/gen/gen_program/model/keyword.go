// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 10:32
// @Software: GoLand

package model

type KeyWord struct {
	Replace   map[string]string `json:"replace"`
	IgnoreDir []string          `json:"ignore_dir"`
}
