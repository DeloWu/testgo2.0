package models
//用例集表
type Suite struct {
    BaseModel `bson:",inline"`
    SuiteName string `bson:"suiteName" json:"suiteName"`
    SuiteSteps []string `bson:"suiteSteps" json:"suiteSteps"`
    SuiteDesc string `bson:"suiteDesc" json:"suiteDesc"`
    RelativePro []string `bson:"relativePro" json:"relativePro"`
    ConfigBaseUrl string `bson:"configBaseUrl" json:"configBaseUrl"`
    // ConfigVariables e.g.  [{"expected_status_code" : 200}, {...}]
    ConfigVariables interface{} `bson:"configVariables" json:"configVariables"`
    //Variables e.g. {"apiId": [{"expected_status_code": 200}], "apiId2": [{"expected_status_code": 200}]}},
    Variables interface{} `bson:"variables" json:"variables"`
    //Parameters e.g. {"caseId": [{"variables1": ["v1", "v2"]}, {"expected_status_code": [201, 404]}]}
    Parameters interface{} `bson:"parameters" json:"parameters"`
}