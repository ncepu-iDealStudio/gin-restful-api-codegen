// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 19:52
// @Software: GoLand

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type UserConfig struct {
	GenCodeConfig GenCodeConfig `yaml:"genCode"`
}
type GenCodeConfig struct {
	TemplatesPath string `yaml:"templates_path"`
	CodeModelPath string `yaml:"code_model_path"`
	DictPath      string `yaml:"dict_path"`
	ResultPath    string `yaml:"result_path"`
}

func GetUserConfig() (*UserConfig, error) {
	var userConfig UserConfig
	var err error
	file, err := ioutil.ReadFile("configs/user.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &userConfig)
	if err != nil {
		return nil, err
	}
	return &userConfig, nil
}
