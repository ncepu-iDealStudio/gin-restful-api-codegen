// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "time"
)

type SsoUserModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    UserID string `gorm:"column:UserID;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    Account string `gorm:"column:Account;type:varchar(255);not null;" json:"Account" form:"Account"`
    Password string `gorm:"column:Password;type:varchar(255);not null;" json:"Password" form:"Password"`
    OpenID string `gorm:"column:OpenID;type:varchar(50);default:;" json:"OpenID" form:"OpenID"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *SsoUserModel) TableName() string {
    return "sso_user"
}

func (m *SsoUserModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *SsoUserModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *SsoUserModel) GetUserID() string {
    return m.UserID
}

func (m *SsoUserModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *SsoUserModel) GetAccount() string {
    return m.Account
}

func (m *SsoUserModel) SetAccount(Account string) {
    m.Account = Account
}

func (m *SsoUserModel) GetPassword() string {
    return m.Password
}

func (m *SsoUserModel) SetPassword(Password string) {
    m.Password = Password
}

func (m *SsoUserModel) GetOpenID() string {
    return m.OpenID
}

func (m *SsoUserModel) SetOpenID(OpenID string) {
    m.OpenID = OpenID
}

func (m *SsoUserModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *SsoUserModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *SsoUserModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *SsoUserModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *SsoUserModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *SsoUserModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *SsoUserModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *SsoUserModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}
