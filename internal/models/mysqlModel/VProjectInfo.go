// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type VProjectInfoModel struct { 
    AutoID int64 `gorm:"column:AutoID;type:bigint(20);not null;default:0;" json:"AutoID" form:"AutoID"`
    ProjectID string `gorm:"column:ProjectID;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    UserID string `gorm:"column:UserID;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    ProjectName string `gorm:"column:ProjectName;type:varchar(255);not null;" json:"ProjectName" form:"ProjectName"`
    ProjectContext string `gorm:"column:ProjectContext;type:text;not null;" json:"ProjectContext" form:"ProjectContext"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:0000-00-00 00:00:00;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
    ProgramTemplate string `gorm:"column:ProgramTemplate;type:text;not null;" json:"ProgramTemplate" form:"ProgramTemplate"`
    ProgramKeyword string `gorm:"column:ProgramKeyword;type:text;not null;" json:"ProgramKeyword" form:"ProgramKeyword"`
    DBCodeTemplate string `gorm:"column:DBCodeTemplate;type:text;not null;" json:"DBCodeTemplate" form:"DBCodeTemplate"`
    DBCodeKeyword string `gorm:"column:DBCodeKeyword;type:text;not null;" json:"DBCodeKeyword" form:"DBCodeKeyword"`
    DataType string `gorm:"column:DataType;type:text;not null;" json:"DataType" form:"DataType"`
    MakeFile string `gorm:"column:MakeFile;type:text;not null;" json:"MakeFile" form:"MakeFile"`
}

func (m *VProjectInfoModel) TableName() string {
    return "v_project_info"
}

func (m *VProjectInfoModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *VProjectInfoModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *VProjectInfoModel) GetProjectID() string {
    return m.ProjectID
}

func (m *VProjectInfoModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *VProjectInfoModel) GetUserID() string {
    return m.UserID
}

func (m *VProjectInfoModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *VProjectInfoModel) GetProjectName() string {
    return m.ProjectName
}

func (m *VProjectInfoModel) SetProjectName(ProjectName string) {
    m.ProjectName = ProjectName
}

func (m *VProjectInfoModel) GetProjectContext() string {
    return m.ProjectContext
}

func (m *VProjectInfoModel) SetProjectContext(ProjectContext string) {
    m.ProjectContext = ProjectContext
}

func (m *VProjectInfoModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *VProjectInfoModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *VProjectInfoModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *VProjectInfoModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *VProjectInfoModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *VProjectInfoModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *VProjectInfoModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *VProjectInfoModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}

func (m *VProjectInfoModel) GetProgramTemplate() string {
    return m.ProgramTemplate
}

func (m *VProjectInfoModel) SetProgramTemplate(ProgramTemplate string) {
    m.ProgramTemplate = ProgramTemplate
}

func (m *VProjectInfoModel) GetProgramKeyword() string {
    return m.ProgramKeyword
}

func (m *VProjectInfoModel) SetProgramKeyword(ProgramKeyword string) {
    m.ProgramKeyword = ProgramKeyword
}

func (m *VProjectInfoModel) GetDBCodeTemplate() string {
    return m.DBCodeTemplate
}

func (m *VProjectInfoModel) SetDBCodeTemplate(DBCodeTemplate string) {
    m.DBCodeTemplate = DBCodeTemplate
}

func (m *VProjectInfoModel) GetDBCodeKeyword() string {
    return m.DBCodeKeyword
}

func (m *VProjectInfoModel) SetDBCodeKeyword(DBCodeKeyword string) {
    m.DBCodeKeyword = DBCodeKeyword
}

func (m *VProjectInfoModel) GetDataType() string {
    return m.DataType
}

func (m *VProjectInfoModel) SetDataType(DataType string) {
    m.DataType = DataType
}

func (m *VProjectInfoModel) GetMakeFile() string {
    return m.MakeFile
}

func (m *VProjectInfoModel) SetMakeFile(MakeFile string) {
    m.MakeFile = MakeFile
}


func (m *VProjectInfoModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}