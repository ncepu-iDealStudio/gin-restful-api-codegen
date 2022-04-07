// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type ProjectMemberModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    ProjectID string `gorm:"column:ProjectID;primaryKey;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    UserID string `gorm:"column:UserID;primaryKey;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    RoleID string `gorm:"column:RoleID;type:varchar(20);not null;" json:"RoleID" form:"RoleID"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *ProjectMemberModel) TableName() string {
    return "project_member"
}

func (m *ProjectMemberModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *ProjectMemberModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *ProjectMemberModel) GetProjectID() string {
    return m.ProjectID
}

func (m *ProjectMemberModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *ProjectMemberModel) GetUserID() string {
    return m.UserID
}

func (m *ProjectMemberModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *ProjectMemberModel) GetRoleID() string {
    return m.RoleID
}

func (m *ProjectMemberModel) SetRoleID(RoleID string) {
    m.RoleID = RoleID
}

func (m *ProjectMemberModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *ProjectMemberModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *ProjectMemberModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *ProjectMemberModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *ProjectMemberModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *ProjectMemberModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *ProjectMemberModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *ProjectMemberModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *ProjectMemberModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *ProjectMemberModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}