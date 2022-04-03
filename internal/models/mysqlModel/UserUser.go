// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type UserUserModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    UserID string `gorm:"column:UserID;primaryKey;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    Name string `gorm:"column:Name;type:varchar(255);not null;" json:"Name" form:"Name"`
    Phone string `gorm:"column:Phone;type:varchar(20);" json:"Phone" form:"Phone"`
    Email string `gorm:"column:Email;type:varchar(255);" json:"Email" form:"Email"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *UserUserModel) TableName() string {
    return "user_User"
}

func (m *UserUserModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *UserUserModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *UserUserModel) GetUserID() string {
    return m.UserID
}

func (m *UserUserModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *UserUserModel) GetName() string {
    return m.Name
}

func (m *UserUserModel) SetName(Name string) {
    m.Name = Name
}

func (m *UserUserModel) GetPhone() string {
    return m.Phone
}

func (m *UserUserModel) SetPhone(Phone string) {
    m.Phone = Phone
}

func (m *UserUserModel) GetEmail() string {
    return m.Email
}

func (m *UserUserModel) SetEmail(Email string) {
    m.Email = Email
}

func (m *UserUserModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *UserUserModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *UserUserModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *UserUserModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *UserUserModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *UserUserModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *UserUserModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *UserUserModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *UserUserModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}