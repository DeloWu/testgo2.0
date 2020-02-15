package models

//项目表
type Project struct {
    BaseModel
    //项目名称
    ProName string `bson:"proName" json:"proName"`
    //项目描述
    ProDesc string `bson:"proDesc" json:"proDesc"`
}