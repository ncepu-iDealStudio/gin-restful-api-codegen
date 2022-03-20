// coding: utf-8
// @Author : lryself
// @Date : 2022/3/3 18:07
// @Software: GoLand

package parser

type ListParser struct {
	Limit  int    `json:"Limit" form:"Limit" binding:"omitempty,numeric"`
	Offset int    `json:"Offset" form:"Offset" binding:"omitempty,numeric"`
	Order  string `json:"Order" form:"Order"`
}
