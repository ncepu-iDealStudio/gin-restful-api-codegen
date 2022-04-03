// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 23:36
// @Software: GoLand

package ginModels

import (
	"LRYGoCodeGen/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserID    string                 `json:"user_id"`
	UserType  string                 `json:"user_type"`
	OtherInfo map[string]interface{} `json:"other_info"`
}

const (
	Platform = "0"
	User     = "1"
)

type userTempService interface {
	GetModelMap() (map[string]interface{}, error)
	Get() error
	SetUserID(string)
}

func NewUser(userID string, userType string) (UserModel, error) {
	var user UserModel
	user.UserID = userID
	user.UserType = userType

	var userService userTempService
	switch userType {
	case Platform:
		userService = &services.UserPlatformAdminService{}
	case User:
		userService = &services.UserUserService{}
	}
	userService.SetUserID(userID)
	err := userService.Get()
	if err != nil {
		return UserModel{}, err
	}
	user.OtherInfo, err = userService.GetModelMap()
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}

func GetUser(c *gin.Context) (UserModel, error) {
	temp, ok := c.Get("user")
	if !ok {
		return UserModel{}, errors.New("无用户")
	}
	user := temp.(UserModel)
	return user, nil
}

func (m UserModel) Auth(allowRole ...string) bool {
	for _, s := range allowRole {
		if s == m.UserType {
			return true
		}
	}
	return false
}
