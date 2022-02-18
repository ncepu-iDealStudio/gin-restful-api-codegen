// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 23:36
// @Software: GoLand

package ginModels

type UserModel struct {
	UserID     string `json:"user_id"`
	IsPlatUser bool   `json:"is_plat_user"`
	IsAdmin    bool   `json:"is_admin"`
}

func (u UserModel) VerifyAdminRole() bool {
	if u.IsPlatUser {
		return true
	}

	if u.IsAdmin {
		return true
	}
	return false
}
