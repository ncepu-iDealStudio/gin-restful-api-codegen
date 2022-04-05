// coding: utf-8
// @Author : lryself
// @Date : 2022/3/30 17:51
// @Software: GoLand

package sso

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/snowflake"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {
	var Parser struct {
		Account     string `form:"Account" json:"Account" binding:"required"`
		Password    string `form:"Password" json:"Password" binding:"required"`
		NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
	}
	var err error
	//解析参数
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
	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, nil)
		return
	}
	if !user.AuthSelf(loginInfo.UserID) {
		parser.JsonAccessDenied(c, "只能修改自己的密码")
		return
	}
	if !user.IsAdmin() {
		//验证密码
		err = bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), []byte(Parser.Password))
		if err != nil {
			parser.JsonDataError(c, "密码错误", nil)
			return
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		parser.JsonInternalError(c, "", err)
		return
	}

	err = loginInfo.Update(map[string]interface{}{
		"Password": string(hash),
	})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", nil)
}

func Register(c *gin.Context) {
	var Parser struct {
		UserID    string `form:"UserID" json:"UserID"`
		Account   string `form:"Account" json:"Account" binding:"required"`
		Password  string `form:"Password" json:"Password" binding:"required"`
		LoginType string `form:"LoginType" json:"LoginType" binding:"required"`
		UserType  string `form:"UserType" json:"UserType" binding:"required"`
		OtherInfo string `form:"OtherInfo" json:"OtherInfo"`
	}
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	//验证是否有权限注册该用户
	if Parser.LoginType != ginModels.User {
		user, err := ginModels.GetUser(c)
		if err != nil {
			parser.JsonGinUserError(c, err)
			return
		}
		if !user.IsAdmin() {
			parser.JsonAccessDenied(c, "您无权注册此用户！")
			return
		}
	}

	//添加sso账号
	var ssoUser services.SsoUserService
	ssoUser.Account = Parser.Account
	err = ssoUser.Get()
	if err != nil {
		if err.Error() != "record not found" {
			parser.JsonDBError(c, "", err)
			return
		}
		ssoUser.UserID = snowflake.GetSnowflakeID()
		//生成加密密码
		hash, err := bcrypt.GenerateFromPassword([]byte(Parser.Password), bcrypt.DefaultCost)
		if err != nil {
			parser.JsonInternalError(c, "", err)
			return
		}

		ssoUser.Account = Parser.Account
		ssoUser.Password = string(hash)
		ssoUser.OtherInfo = Parser.OtherInfo

		err = ssoUser.Add()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	}

	var userService userInterface
	switch Parser.LoginType {
	case ginModels.Platform:
		userService = &services.UserPlatformAdminService{}
	case ginModels.User:
		userService = &services.UserUserService{}
	default:
		parser.JsonParameterIllegal(c, "登录的用户类型不合法！", errors.New("注册的用户类型不合法！"))
		return
	}
	userService.SetUserID(ssoUser.UserID)
	err = userService.Get()
	if err != nil {
		if err.Error() != "record not found" {
			parser.JsonDBError(c, "", err)
			return
		}
	}

	//开始注册
	switch Parser.UserType {
	case ginModels.Platform:
		var userInfoService services.UserPlatformAdminService
		err = c.ShouldBind(&userInfoService)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		userService = &userInfoService
	case ginModels.User:
		var userInfoService services.UserUserService
		err = c.ShouldBind(&userInfoService)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		userService = &userInfoService
	default:
		parser.JsonParameterIllegal(c, "登录的用户类型不合法！", errors.New("注册的用户类型不合法！"))
		return
	}

	userService.SetUserID(ssoUser.UserID)
	err = userService.Add()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	backMap, err := userService.GetModelMap()
	if err != nil {
		parser.JsonInternalError(c, "", err)
	}
	parser.JsonOK(c, "", backMap)
}
