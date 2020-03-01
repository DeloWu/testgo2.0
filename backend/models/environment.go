package models

//环境表
type Environment struct {
	BaseModel   `bson:",inline"`
	EnvName     string   `bson:"envName" json:"envName"`
	EnvIp       string   `bson:"envIp" json:"envIp"`
	EnvPort     string   `bson:"envPort" json:"envPort"`
	EnvDesc     string   `bson:"envDesc" json:"envDesc"`
	RelativePro []string `bson:"relativePro" json:"relativePro"`
}
