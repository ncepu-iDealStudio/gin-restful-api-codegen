// coding: utf-8
// @Author : lryself
// @Date : 2022/2/25 17:22
// @Software: GoLand

package str

import "strings"

func LineToUpCamel(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	return strings.ReplaceAll(s, " ", "")
}

func LineToLowCamel(s string) string {
	s = LineToUpCamel(s)
	s = strings.ToUpper(string([]rune(s)[0])) + string([]rune(s)[1:])
	return s
}
