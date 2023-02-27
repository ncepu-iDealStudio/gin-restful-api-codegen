// Package keywords
/*
@Coding : utf-8
@Time : 2023/2/22 16:31
@Author : yizhigopher
@Software : GoLand
*/
package keywords

func In(target interface{}) bool {
	for _, elem := range keyWords {
		if target == elem {
			return true
		}
	}
	return false
}
