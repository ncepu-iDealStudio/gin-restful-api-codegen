// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 23:36
// @Software: GoLand

package ginModels

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserID    string                 `json:"user_id"`
	UserType  string                 `json:"user_type"`
	OtherInfo map[string]interface{} `json:"other_info"`
}

const (
	PlatformUser = "0" // 平台用户
)

type userTempService interface {
	GetModelMap() (map[string]interface{}, error)
	Get() error
	SetUserID(string)
}

func GetUser(c *gin.Context) (UserModel, error) {
	temp, ok := c.Get("user")
	if !ok {
		return UserModel{}, errors.New("无用户")
	}
	user := temp.(UserModel)
	return user, nil
}

func NewUser(userID string, userType string) (UserModel, error) {
	var err error
	user := UserModel{UserID: "", UserType: "", OtherInfo: map[string]interface{}{}}
	user.UserID = userID
	user.UserType = userType

	//var userService userTempService
	//switch userType {
	//case PlatformUser:
	//	userService = &services.PlatformUserInfoService{}
	//}
	//userService.SetUserID(userID)
	//err := userService.Get()
	//if err != nil {
	//	return UserModel{}, err
	//}
	//user.OtherInfo, err = userService.GetModelMap()
	//if err != nil {
	//	return UserModel{}, err
	//}
	return user, err
}

//func (u UserModel) IsPlatform() bool {
//	if u.UserType == Platform {
//		return true
//	}
//	return false
//}
