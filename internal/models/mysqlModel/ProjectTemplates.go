// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type ProjectTemplatesModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    ProjectID string `gorm:"column:ProjectID;primaryKey;type:varchar(20);not null;" json:"ProjectID" form:"ProjectID"`
    ProgramTemplate string `gorm:"column:ProgramTemplate;type:text;not null;" json:"ProgramTemplate" form:"ProgramTemplate"`
    ProgramKeyword string `gorm:"column:ProgramKeyword;type:text;not null;" json:"ProgramKeyword" form:"ProgramKeyword"`
    DBCodeTemplate string `gorm:"column:DBCodeTemplate;type:text;not null;" json:"DBCodeTemplate" form:"DBCodeTemplate"`
    DBCodeKeyword string `gorm:"column:DBCodeKeyword;type:text;not null;" json:"DBCodeKeyword" form:"DBCodeKeyword"`
    DataType string `gorm:"column:DataType;type:text;not null;" json:"DataType" form:"DataType"`
    MakeFile string `gorm:"column:MakeFile;type:text;not null;" json:"MakeFile" form:"MakeFile"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *ProjectTemplatesModel) TableName() string {
    return "project_templates"
}

func (m *ProjectTemplatesModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *ProjectTemplatesModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *ProjectTemplatesModel) GetProjectID() string {
    return m.ProjectID
}

func (m *ProjectTemplatesModel) SetProjectID(ProjectID string) {
    m.ProjectID = ProjectID
}

func (m *ProjectTemplatesModel) GetProgramTemplate() string {
    return m.ProgramTemplate
}

func (m *ProjectTemplatesModel) SetProgramTemplate(ProgramTemplate string) {
    m.ProgramTemplate = ProgramTemplate
}

func (m *ProjectTemplatesModel) GetProgramKeyword() string {
    return m.ProgramKeyword
}

func (m *ProjectTemplatesModel) SetProgramKeyword(ProgramKeyword string) {
    m.ProgramKeyword = ProgramKeyword
}

func (m *ProjectTemplatesModel) GetDBCodeTemplate() string {
    return m.DBCodeTemplate
}

func (m *ProjectTemplatesModel) SetDBCodeTemplate(DBCodeTemplate string) {
    m.DBCodeTemplate = DBCodeTemplate
}

func (m *ProjectTemplatesModel) GetDBCodeKeyword() string {
    return m.DBCodeKeyword
}

func (m *ProjectTemplatesModel) SetDBCodeKeyword(DBCodeKeyword string) {
    m.DBCodeKeyword = DBCodeKeyword
}

func (m *ProjectTemplatesModel) GetDataType() string {
    return m.DataType
}

func (m *ProjectTemplatesModel) SetDataType(DataType string) {
    m.DataType = DataType
}

func (m *ProjectTemplatesModel) GetMakeFile() string {
    return m.MakeFile
}

func (m *ProjectTemplatesModel) SetMakeFile(MakeFile string) {
    m.MakeFile = MakeFile
}

func (m *ProjectTemplatesModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *ProjectTemplatesModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *ProjectTemplatesModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *ProjectTemplatesModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *ProjectTemplatesModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *ProjectTemplatesModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *ProjectTemplatesModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *ProjectTemplatesModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *ProjectTemplatesModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *ProjectTemplatesModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}