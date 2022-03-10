// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 10:32
// @Software: GoLand

package model

type KeyWord struct {
	Include   map[string]string `json:"include"`
	IgnoreDir []string          `json:"ignore_dir"`
}
