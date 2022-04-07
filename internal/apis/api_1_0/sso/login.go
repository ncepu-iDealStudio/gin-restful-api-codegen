// coding: utf-8
// @Author : lryself
// @Date : 2022/3/30 17:48
// @Software: GoLand

package sso

import (
	"LRYGoCodeGen/internal/globals/codes"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/jwt"
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/rsa"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"gitee.com/lryself/go-utils/email"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"net/http"
	"time"
)

type userInterface interface {
	Add() error
	Get() error
	SetUserID(string)
	SetName(string)
	GetModelMap() (map[string]interface{}, error)
}

func Login(c *gin.Context) {
	var err error
	var LoginParser struct {
		LoginMethod string `form:"LoginMethod" json:"LoginMethod" binding:"required"`
		Account     string `form:"Account" json:"Account" binding:"required"`
		LoginType   string `form:"LoginType" json:"LoginType" binding:"required"`
	}
	err = c.ShouldBind(&LoginParser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var loginInfo services.SsoUserService
	loginInfo.Account = LoginParser.Account
	err = loginInfo.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	//校验密码
	switch LoginParser.LoginMethod {
	case "password":
		var Parser struct {
			Account  string `form:"Account" json:"Account" binding:"required"`
			Password string `form:"Password" json:"Password" binding:"required"`
		}
		//解析参数
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		// RSA解密密码
		var password []byte
		password, err = base64.StdEncoding.DecodeString(Parser.Password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "解码失败！",
				"err":     err,
			})
			return
		}
		RSA := rsa.GetRSAHelper()
		password, err = RSA.Decrypt(password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "解密失败！",
				"err":     err,
			})
			return
		}
		// 校验密码
		err = bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), password)
		if err != nil {
			if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
				c.JSON(http.StatusOK, gin.H{
					"code":    codes.AccessDenied,
					"message": "密码错误！",
				})
				return
			}
			parser.JsonInternalError(c, "", err)
			return
		}
	case "email":
		var Parser struct {
			Account    string `form:"Account" json:"Account" binding:"required,email"`
			VerifyCode string `form:"VerifyCode" json:"VerifyCode" binding:"required"`
		}
		//解析参数
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		// 查询验证码
		redisManager := database.GetRedisManager()
		stringCmd := redisManager.Get("verify_" + Parser.Account)
		if stringCmd.Err() != nil {
			parser.JsonDBError(c, "未找到该账号验证码！", stringCmd.Err())
			return
		}

		verifyCode := stringCmd.Val()
		if Parser.VerifyCode != verifyCode {
			parser.JsonDataError(c, "验证码错误！", nil)
			return
		}
	default:
		parser.JsonParameterIllegal(c, "登录方式不合法", errors.New("登录方式不合法"))
		return
	}
	//校验账号是否有该LoginType对应的角色
	var userService userInterface
	switch LoginParser.LoginType {
	case ginModels.Platform:
		userService = &services.UserPlatformAdminService{}
	case ginModels.User:
		userService = &services.UserUserService{}
	default:
		parser.JsonParameterIllegal(c, "登录的用户类型不合法！", errors.New("登录的用户类型不合法！"))
		return
	}
	userService.SetUserID(loginInfo.UserID)
	err = userService.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// 生成Token
	claims := jwt.JWTClaims{UserID: loginInfo.UserID, UserType: LoginParser.LoginType}
	token, err := claims.MakeToken(viper.GetInt("system.TokenExpireTime"), []byte(viper.GetString("system.Secret")))
	if err != nil {
		parser.JsonInternalError(c, "", err)
		return
	}
	// 存入redis
	redisManager := database.GetRedisManager()
	err = redisManager.Set("Token_"+loginInfo.UserID, token, time.Duration(viper.GetInt("system.RedisExpireTime"))*time.Second).Err()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	var err error
	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
	}
	redisManager := database.GetRedisManager()
	err = redisManager.Del("Token_" + user.UserID).Err()
	if err != nil {
		parser.JsonInternalError(c, "", err)
		return
	}
	parser.JsonOK(c, "", nil)
	return
}

func MakeEmailVerifyCode(c *gin.Context) {
	var Parser struct {
		Account string `form:"Account" json:"Account" binding:"required,email"`
	}
	var err error
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var loginInfo services.SsoUserService
	loginInfo.Account = Parser.Account
	err = loginInfo.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	redisManager := database.GetRedisManager()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	stringCmd := redisManager.Get("verify_" + Parser.Account)
	err = stringCmd.Err()

	var verifyCode string
	if err != nil {
		if err == redis.Nil {
			verifyCode = createRandomString(8)
			err = redisManager.Set("verify_"+Parser.Account, verifyCode, time.Duration(300)*time.Second).Err()
			if err != nil {
				parser.JsonDBError(c, "", err)
				return
			}
		} else {
			parser.JsonDBError(c, "", err)
			return
		}
	} else {
		verifyCode = stringCmd.Val()
	}

	Email := email.SMTPClient{
		SMTPHost: viper.GetString("smtp.Host"),
		SMTPPort: viper.GetString("smtp.Port"),
		SMTPUser: viper.GetString("smtp.User"),
		SMTPPass: viper.GetString("smtp.Pass"),
	}

	err = Email.SMTPSendEmail(
		"登录中心",
		Parser.Account,
		"【登录中心】您的验证码是"+verifyCode,
		"plain",
		"【登录中心】您的验证码是"+verifyCode+"。验证码在5分钟内有效，如果不是您的邮件，请忽略此邮件。",
	)

	if err != nil {
		parser.JsonInternalError(c, "", err)
		return
	}

	parser.JsonOK(c, "", nil)
	return
}

func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
