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
	ConfigHeaders interface{} `bson:"configHeaders" json:"configHeaders"`
	// 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
	ConfigVariables interface{} `bson:"configVariables" json:"configVariables"`
	// ConfigOutput e.g.  ["session_token", ...]
	ConfigOutput interface{} `bson:"configOutput" json:"configOutput"`
	//ConfigVerify 默认为false
	ConfigVerify bool `bson:"configVerify" json:"configVerify"`
	//Extracts/ Validates/ Variables 仅需要给api步骤添加额外参数
	// 格式为: [{"id": "", "keyName":"", "value":""}, {, , }...]
	StepExtracts interface{} `bson:"stepExtracts" json:"stepExtracts"`
	// 格式为: [{"id": "", "keyName":"", "comparator":"", "value":""}, {, , }...]
	StepValidates interface{} `bson:"stepValidates" json:"stepValidates"`
	// 格式为: [{"id": "", "keyName":"", "valueType":"", "value":""}, {, , }...]
	StepVariables interface{} `bson:"stepVariables" json:"stepVariables"`
	//SetupHooks e.g. [{"value":"${setup_hook_prepare_kwargs($request)}"},{"value":"${teardown_hook_sleep_N_secs($response, n_secs)}"}],
	SetupHooks    interface{} `bson:"setupHooks" json:"setupHooks"`
	TeardownHooks interface{} `bson:"teardownHooks" json:"teardownHooks"`
}
