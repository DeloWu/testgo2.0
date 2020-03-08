package models

//用例表
type Case struct {
	BaseModel `bson:",inline"`
	CaseName  string `bson:"caseName" json:"caseName"`
	// CaseSteps  e.g. [{"foo_obj_id": "api"}, {"var_obj_id": "testcase"}, {...}]
	CaseSteps     interface{} `bson:"caseSteps" json:"caseSteps"`
	RelativePro   []string    `bson:"relativePro" json:"relativePro"`
	CaseDesc      string      `bson:"caseDesc" json:"caseDesc"`
	ConfigBaseUrl string      `bson:"configBaseUrl" json:"configBaseUrl"`
	// 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
	ConfigVariables interface{} `bson:"configVariables" json:"configVariables"`
	// ConfigOutput e.g.  ["session_token", ...]
	ConfigOutput []string `bson:"configOutput" json:"configOutput"`
	//ConfigVerify 默认为false
	ConfigVerify bool `bson:"configVerify" json:"configVerify"`
	// 格式为: [{"keyName":"", "value":""}, {, , }...]
	Extract interface{} `bson:"extract" json:"extract"`
	// 格式为: [{"keyName":"", "comparator":"", "value":""}], {, , }...
	Validate interface{} `bson:"validate" json:"validate"`
	// 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
	Variables interface{} `bson:"variables" json:"variables"`
	//SetupHooks e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"],
	SetupHooks []string `bson:"setupHooks" json:"setupHooks"`
	//TeardownHooks e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"]
	TeardownHooks []string `bson:"teardownHooks" json:"teardownHooks"`
}
