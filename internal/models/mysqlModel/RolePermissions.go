// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type RolePermissionsModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    PermissionID string `gorm:"column:PermissionID;primaryKey;type:varchar(20);not null;" json:"PermissionID" form:"PermissionID"`
    PermissionName string `gorm:"column:PermissionName;type:varchar(255);" json:"PermissionName" form:"PermissionName"`
    PermissionContext string `gorm:"column:PermissionContext;type:text;" json:"PermissionContext" form:"PermissionContext"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *RolePermissionsModel) TableName() string {
    return "role_permissions"
}

func (m *RolePermissionsModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *RolePermissionsModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *RolePermissionsModel) GetPermissionID() string {
    return m.PermissionID
}

func (m *RolePermissionsModel) SetPermissionID(PermissionID string) {
    m.PermissionID = PermissionID
}

func (m *RolePermissionsModel) GetPermissionName() string {
    return m.PermissionName
}

func (m *RolePermissionsModel) SetPermissionName(PermissionName string) {
    m.PermissionName = PermissionName
}

func (m *RolePermissionsModel) GetPermissionContext() string {
    return m.PermissionContext
}

func (m *RolePermissionsModel) SetPermissionContext(PermissionContext string) {
    m.PermissionContext = PermissionContext
}

func (m *RolePermissionsModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *RolePermissionsModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *RolePermissionsModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *RolePermissionsModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *RolePermissionsModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *RolePermissionsModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *RolePermissionsModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *RolePermissionsModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *RolePermissionsModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *RolePermissionsModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}