// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type RolePoolModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    RoleID string `gorm:"column:RoleID;primaryKey;type:varchar(20);not null;" json:"RoleID" form:"RoleID"`
    RoleName string `gorm:"column:RoleName;type:varchar(255);not null;" json:"RoleName" form:"RoleName"`
    PermissionList string `gorm:"column:PermissionList;type:text;" json:"PermissionList" form:"PermissionList"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *RolePoolModel) TableName() string {
    return "role_pool"
}

func (m *RolePoolModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *RolePoolModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *RolePoolModel) GetRoleID() string {
    return m.RoleID
}

func (m *RolePoolModel) SetRoleID(RoleID string) {
    m.RoleID = RoleID
}

func (m *RolePoolModel) GetRoleName() string {
    return m.RoleName
}

func (m *RolePoolModel) SetRoleName(RoleName string) {
    m.RoleName = RoleName
}

func (m *RolePoolModel) GetPermissionList() string {
    return m.PermissionList
}

func (m *RolePoolModel) SetPermissionList(PermissionList string) {
    m.PermissionList = PermissionList
}

func (m *RolePoolModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *RolePoolModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *RolePoolModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *RolePoolModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *RolePoolModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *RolePoolModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *RolePoolModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *RolePoolModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *RolePoolModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *RolePoolModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}