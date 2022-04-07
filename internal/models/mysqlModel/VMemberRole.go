// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type VMemberRoleModel struct { 
    AutoID int64 `gorm:"column:AutoID;type:bigint(20);not null;default:0;" json:"AutoID" form:"AutoID"`
    ProjectID string `gorm:"column:ProjectID;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    UserID string `gorm:"column:UserID;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    RoleID string `gorm:"column:RoleID;type:varchar(20);not null;" json:"RoleID" form:"RoleID"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
    RoleName string `gorm:"column:RoleName;type:varchar(255);not null;" json:"RoleName" form:"RoleName"`
    PermissionList string `gorm:"column:PermissionList;type:text;" json:"PermissionList" form:"PermissionList"`
    Name string `gorm:"column:Name;type:varchar(255);not null;" json:"Name" form:"Name"`
    Phone string `gorm:"column:Phone;type:varchar(20);" json:"Phone" form:"Phone"`
    Email string `gorm:"column:Email;type:varchar(255);" json:"Email" form:"Email"`
}

func (m *VMemberRoleModel) TableName() string {
    return "v_member_role"
}

func (m *VMemberRoleModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *VMemberRoleModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *VMemberRoleModel) GetProjectID() string {
    return m.ProjectID
}

func (m *VMemberRoleModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *VMemberRoleModel) GetUserID() string {
    return m.UserID
}

func (m *VMemberRoleModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *VMemberRoleModel) GetRoleID() string {
    return m.RoleID
}

func (m *VMemberRoleModel) SetRoleID(RoleID string) {
    m.RoleID = RoleID
}

func (m *VMemberRoleModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *VMemberRoleModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *VMemberRoleModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *VMemberRoleModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *VMemberRoleModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *VMemberRoleModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *VMemberRoleModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *VMemberRoleModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}

func (m *VMemberRoleModel) GetRoleName() string {
    return m.RoleName
}

func (m *VMemberRoleModel) SetRoleName(RoleName string) {
    m.RoleName = RoleName
}

func (m *VMemberRoleModel) GetPermissionList() string {
    return m.PermissionList
}

func (m *VMemberRoleModel) SetPermissionList(PermissionList string) {
    m.PermissionList = PermissionList
}

func (m *VMemberRoleModel) GetName() string {
    return m.Name
}

func (m *VMemberRoleModel) SetName(Name string) {
    m.Name = Name
}

func (m *VMemberRoleModel) GetPhone() string {
    return m.Phone
}

func (m *VMemberRoleModel) SetPhone(Phone string) {
    m.Phone = Phone
}

func (m *VMemberRoleModel) GetEmail() string {
    return m.Email
}

func (m *VMemberRoleModel) SetEmail(Email string) {
    m.Email = Email
}


func (m *VMemberRoleModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *VMemberRoleModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}