package models

type Report struct {
    BaseModel `bson:",inline"`
    ReportCreator string `bson:"reportCreator" json:"reportCreator"`
    ReportPath string `bson:"reportPath" json:"reportPath"`

}
