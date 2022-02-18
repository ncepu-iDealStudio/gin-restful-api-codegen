// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 1:48
// @Software: GoLand

package rsa

import (
	"github.com/spf13/viper"
	"tem_go_project/internal/utils/rsa"
)

func GetRSAHelper() rsa.RSA {
	return rsa.RSA{
		PublicKeyPath:  viper.GetString("system.RSAPublic"),
		PrivateKeyPath: viper.GetString("system.RSAPrivate"),
	}
}
