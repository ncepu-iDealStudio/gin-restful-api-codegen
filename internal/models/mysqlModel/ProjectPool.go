// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type ProjectPoolModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    ProjectID string `gorm:"column:ProjectID;primaryKey;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    UserID string `gorm:"column:UserID;primaryKey;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    ProjectName string `gorm:"column:ProjectName;type:varchar(255);not null;" json:"ProjectName" form:"ProjectName"`
    ProjectType string `gorm:"column:ProjectType;type:varchar(20);not null;" json:"ProjectType" form:"ProjectType"`
    ProjectContext string `gorm:"column:ProjectContext;type:text;" json:"ProjectContext" form:"ProjectContext"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *ProjectPoolModel) TableName() string {
    return "project_pool"
}

func (m *ProjectPoolModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *ProjectPoolModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *ProjectPoolModel) GetProjectID() string {
    return m.ProjectID
}

func (m *ProjectPoolModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *ProjectPoolModel) GetUserID() string {
    return m.UserID
}

func (m *ProjectPoolModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *ProjectPoolModel) GetProjectName() string {
    return m.ProjectName
}

func (m *ProjectPoolModel) SetProjectName(ProjectName string) {
    m.ProjectName = ProjectName
}

func (m *ProjectPoolModel) GetProjectType() string {
    return m.ProjectType
}

func (m *ProjectPoolModel) SetProjectType(ProjectType string) {
    m.ProjectType = ProjectType
}

func (m *ProjectPoolModel) GetProjectContext() string {
    return m.ProjectContext
}

func (m *ProjectPoolModel) SetProjectContext(ProjectContext string) {
    m.ProjectContext = ProjectContext
}

func (m *ProjectPoolModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *ProjectPoolModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *ProjectPoolModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *ProjectPoolModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *ProjectPoolModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *ProjectPoolModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *ProjectPoolModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *ProjectPoolModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *ProjectPoolModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}