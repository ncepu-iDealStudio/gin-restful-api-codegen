// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type UserStuffUserModel struct { 
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

func (m *UserStuffUserModel) TableName() string {
    return "user_stuffUser"
}

func (m *UserStuffUserModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *UserStuffUserModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *UserStuffUserModel) GetUserID() string {
    return m.UserID
}

func (m *UserStuffUserModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *UserStuffUserModel) GetName() string {
    return m.Name
}

func (m *UserStuffUserModel) SetName(Name string) {
    m.Name = Name
}

func (m *UserStuffUserModel) GetPhone() string {
    return m.Phone
}

func (m *UserStuffUserModel) SetPhone(Phone string) {
    m.Phone = Phone
}

func (m *UserStuffUserModel) GetEmail() string {
    return m.Email
}

func (m *UserStuffUserModel) SetEmail(Email string) {
    m.Email = Email
}

func (m *UserStuffUserModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *UserStuffUserModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *UserStuffUserModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *UserStuffUserModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *UserStuffUserModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *UserStuffUserModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *UserStuffUserModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *UserStuffUserModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *UserStuffUserModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}