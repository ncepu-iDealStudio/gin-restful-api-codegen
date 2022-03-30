// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 23:36
// @Software: GoLand

package ginModels

import "LRYGoCodeGen/internal/services"

type UserModel struct {
	UserID    string                 `json:"user_id"`
	UserType  string                 `json:"user_type"`
	RoleID    bool                   `json:"role_id"`
	OtherInfo map[string]interface{} `json:"other_info"`
}

const (
	Platform   = "0"
	StuffAdmin = "1"
	StuffUser  = "2"
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
	case StuffUser:
		userService = &services.UserStuffUserService{}
	case StuffAdmin:
		userService = &services.UserStuffAdminService{}
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
