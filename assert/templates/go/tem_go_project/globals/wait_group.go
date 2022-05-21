// coding: utf-8
// @Author : lryself
// @Date : 2022/5/14 17:22
// @Software: GoLand

package globals

import "sync"

var waitGroup sync.WaitGroup

func GetWatGroup() *sync.WaitGroup {
	return &waitGroup
}
