// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type UserPlatformAdminModel struct { 
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

func (m *UserPlatformAdminModel) TableName() string {
    return "user_platformAdmin"
}

func (m *UserPlatformAdminModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *UserPlatformAdminModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *UserPlatformAdminModel) GetUserID() string {
    return m.UserID
}

func (m *UserPlatformAdminModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *UserPlatformAdminModel) GetName() string {
    return m.Name
}

func (m *UserPlatformAdminModel) SetName(Name string) {
    m.Name = Name
}

func (m *UserPlatformAdminModel) GetPhone() string {
    return m.Phone
}

func (m *UserPlatformAdminModel) SetPhone(Phone string) {
    m.Phone = Phone
}

func (m *UserPlatformAdminModel) GetEmail() string {
    return m.Email
}

func (m *UserPlatformAdminModel) SetEmail(Email string) {
    m.Email = Email
}

func (m *UserPlatformAdminModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *UserPlatformAdminModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *UserPlatformAdminModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *UserPlatformAdminModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *UserPlatformAdminModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *UserPlatformAdminModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *UserPlatformAdminModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *UserPlatformAdminModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *UserPlatformAdminModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *UserPlatformAdminModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}