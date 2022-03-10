// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:32
// @Software: GoLand

package models

import (
	"fmt"
	"regexp"
)

type TypeDict struct {
	Accurate map[string]string `json:"accurate"`
	Fuzzy    map[string]string `json:"fuzzy"`
}

func (t *TypeDict) GetGoType(name string) string {
	// Precise matching first.先精确匹配
	if v, ok := t.Accurate[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range t.Fuzzy {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way", name))
}

type structModel struct {
	Header string `json:"header,omitempty"`
	Row    string `json:"row,omitempty"`
	Footer string `json:"footer,omitempty"`
}
type CodeModel struct {
	StaticDict map[string]string `json:"static_dict,omitempty"`
	Filepath   string            `json:"filepath,omitempty"`
	FileHeader string            `json:"file_header,omitempty"`
	Import     structModel       `json:"import,omitempty"`
	Struct     structModel       `json:"struct,omitempty"`
	Methods    map[string]string `json:"methods,omitempty"`
}
