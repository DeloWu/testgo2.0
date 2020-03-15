package models

//用例集表
type Suite struct {
	BaseModel       `bson:",inline"`
	SuiteName       string      `bson:"suiteName" json:"suiteName"`
	SuiteSteps      []string    `bson:"suiteSteps" json:"suiteSteps"`
	SuiteDesc       string      `bson:"suiteDesc" json:"suiteDesc"`
	RelativePro     []string    `bson:"relativePro" json:"relativePro"`
	ConfigBaseUrl   string      `bson:"configBaseUrl" json:"configBaseUrl"`
	ConfigVariables interface{} `bson:"configVariables" json:"configVariables"`
	StepVariables   interface{} `bson:"stepVariables" json:"stepVariables"`
	StepParameters  interface{} `bson:"stepParameters" json:"stepParameters"`
}
