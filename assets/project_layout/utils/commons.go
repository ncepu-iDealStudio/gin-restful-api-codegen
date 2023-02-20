// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 19:50
// @Software: GoLand

package utils

import (
	"os"
	"sync"
)

var waitGroup sync.WaitGroup

func GetWaitGroup() *sync.WaitGroup {
	return &waitGroup
}

func PathExists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
