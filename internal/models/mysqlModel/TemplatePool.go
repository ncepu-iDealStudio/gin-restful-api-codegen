// coding: utf-8
// @Author : lryself
// @Software: GoLand

package mysqlModel

import (
    "LRYGoCodeGen/internal/utils"
    "gitee.com/lryself/go-utils/structs"
    "time"
)

type TemplatePoolModel struct { 
    AutoID int64 `gorm:"column:AutoID;primaryKey;type:bigint(20);not null;" json:"AutoID" form:"AutoID"`
    TemplateID string `gorm:"column:TemplateID;primaryKey;type:varchar(20);not null;" json:"TemplateID" form:"TemplateID"`
    UserID string `gorm:"column:UserID;primaryKey;type:varchar(20);not null;" json:"UserID" form:"UserID"`
    TemplateName string `gorm:"column:TemplateName;type:varchar(255);not null;" json:"TemplateName" form:"TemplateName"`
    TemplateType int8 `gorm:"column:TemplateType;type:tinyint(2);not null;" json:"TemplateType" form:"TemplateType"`
    StoreType int8 `gorm:"column:StoreType;type:tinyint(2);not null;" json:"StoreType" form:"StoreType"`
    StorePath string `gorm:"column:StorePath;type:varchar(255);not null;" json:"StorePath" form:"StorePath"`
    IsPublic bool `gorm:"column:IsPublic;type:tinyint(1);not null;default:0;" json:"IsPublic" form:"IsPublic"`
    OtherInfo string `gorm:"column:OtherInfo;type:text;" json:"OtherInfo" form:"OtherInfo"`
    Createtime time.Time `gorm:"column:Createtime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Createtime" form:"Createtime"`
    Updatetime time.Time `gorm:"column:Updatetime;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"Updatetime" form:"Updatetime"`
    IsDeleted bool `gorm:"column:IsDeleted;type:tinyint(1);not null;default:0;" json:"IsDeleted" form:"IsDeleted"`
}

func (m *TemplatePoolModel) TableName() string {
    return "template_pool"
}

func (m *TemplatePoolModel) GetAutoID() int64 {
    return m.AutoID
}

func (m *TemplatePoolModel) SetAutoID(AutoID int64) {
    m.AutoID = AutoID
}

func (m *TemplatePoolModel) GetTemplateID() string {
    return m.TemplateID
}

func (m *TemplatePoolModel) SetTemplateID(TemplateID string) {
    m.TemplateID = TemplateID
}

func (m *TemplatePoolModel) GetUserID() string {
    return m.UserID
}

func (m *TemplatePoolModel) SetUserID(UserID string) {
    m.UserID = UserID
}

func (m *TemplatePoolModel) GetTemplateName() string {
    return m.TemplateName
}

func (m *TemplatePoolModel) SetTemplateName(TemplateName string) {
    m.TemplateName = TemplateName
}

func (m *TemplatePoolModel) GetTemplateType() int8 {
    return m.TemplateType
}

func (m *TemplatePoolModel) SetTemplateType(TemplateType int8) {
    m.TemplateType = TemplateType
}

func (m *TemplatePoolModel) GetStoreType() int8 {
    return m.StoreType
}

func (m *TemplatePoolModel) SetStoreType(StoreType int8) {
    m.StoreType = StoreType
}

func (m *TemplatePoolModel) GetStorePath() string {
    return m.StorePath
}

func (m *TemplatePoolModel) SetStorePath(StorePath string) {
    m.StorePath = StorePath
}

func (m *TemplatePoolModel) GetIsPublic() bool {
    return m.IsPublic
}

func (m *TemplatePoolModel) SetIsPublic(IsPublic bool) {
    m.IsPublic = IsPublic
}

func (m *TemplatePoolModel) GetOtherInfo() string {
    return m.OtherInfo
}

func (m *TemplatePoolModel) SetOtherInfo(OtherInfo string) {
    m.OtherInfo = OtherInfo
}

func (m *TemplatePoolModel) GetCreatetime() time.Time {
    return m.Createtime
}

func (m *TemplatePoolModel) SetCreatetime(Createtime time.Time) {
    m.Createtime = Createtime
}

func (m *TemplatePoolModel) GetUpdatetime() time.Time {
    return m.Updatetime
}

func (m *TemplatePoolModel) SetUpdatetime(Updatetime time.Time) {
    m.Updatetime = Updatetime
}

func (m *TemplatePoolModel) GetIsDeleted() bool {
    return m.IsDeleted
}

func (m *TemplatePoolModel) SetIsDeleted(IsDeleted bool) {
    m.IsDeleted = IsDeleted
}


func (m *TemplatePoolModel) GetModelMap() (map[string]interface{}, error) {
    return structs.StructToMap(m, "json")
}

func (m *TemplatePoolModel) Assign(in interface{}) {
    utils.StructAssign(m, in, "json")
}