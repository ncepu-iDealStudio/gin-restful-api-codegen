// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type RolePremissionsModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    PermissionID string `gorm:"column:PermissionID;primaryKey;type:varchar(20);not null;" json:"PermissionID" form:"PermissionID"`
    PermissionName string `gorm:"column:PermissionName;type:varchar(255);" json:"PermissionName" form:"PermissionName"`
    PermissionContext string `gorm:"column:PermissionContext;type:text;" json:"PermissionContext" form:"PermissionContext"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *RolePremissionsModel) TableName() string {
    return "role_premissions"
}

func (m *RolePremissionsModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *RolePremissionsModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *RolePremissionsModel) GetPermissionID() string {
    return m.PermissionID
}

func (m *RolePremissionsModel) SetPermissionID(PermissionID string) {
    m.PermissionID = PermissionID
}

func (m *RolePremissionsModel) GetPermissionName() string {
    return m.PermissionName
}

func (m *RolePremissionsModel) SetPermissionName(PermissionName string) {
    m.PermissionName = PermissionName
}

func (m *RolePremissionsModel) GetPermissionContext() string {
    return m.PermissionContext
}

func (m *RolePremissionsModel) SetPermissionContext(PermissionContext string) {
    m.PermissionContext = PermissionContext
}

func (m *RolePremissionsModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *RolePremissionsModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *RolePremissionsModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *RolePremissionsModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *RolePremissionsModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *RolePremissionsModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *RolePremissionsModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *RolePremissionsModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *RolePremissionsModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}