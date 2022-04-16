// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type VProjectMemberModel struct { 
    ProjectID string `gorm:"column:ProjectID;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    UserID string `gorm:"column:UserID;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    ProjectName string `gorm:"column:ProjectName;type:varchar(255);not null;" json:"ProjectName" form:"ProjectName"`
    ProjectContext string `gorm:"column:ProjectContext;type:text;not null;" json:"ProjectContext" form:"ProjectContext"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Updatetime" form:"Updatetime"`
    MemberID string `gorm:"column:MemberID;type:varchar(20);not null;" json:"MemberID" form:"MemberID"`
    RoleID string `gorm:"column:RoleID;type:varchar(20);not null;" json:"RoleID" form:"RoleID"`
}

func (m *VProjectMemberModel) TableName() string {
    return "v_project_member"
}

func (m *VProjectMemberModel) GetProjectID() string {
    return m.ProjectID
}

func (m *VProjectMemberModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *VProjectMemberModel) GetUserID() string {
    return m.UserID
}

func (m *VProjectMemberModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *VProjectMemberModel) GetProjectName() string {
    return m.ProjectName
}

func (m *VProjectMemberModel) SetProjectName(ProjectName string) {
    m.ProjectName = ProjectName
}

func (m *VProjectMemberModel) GetProjectContext() string {
    return m.ProjectContext
}

func (m *VProjectMemberModel) SetProjectContext(ProjectContext string) {
    m.ProjectContext = ProjectContext
}

func (m *VProjectMemberModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *VProjectMemberModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *VProjectMemberModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *VProjectMemberModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *VProjectMemberModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *VProjectMemberModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *VProjectMemberModel) GetMemberID() string {
    return m.MemberID
}

func (m *VProjectMemberModel) SetMemberID(MemberID string) {
    m.MemberID = MemberID
}

func (m *VProjectMemberModel) GetRoleID() string {
    return m.RoleID
}

func (m *VProjectMemberModel) SetRoleID(RoleID string) {
    m.RoleID = RoleID
}


func (m *VProjectMemberModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *VProjectMemberModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}