package models
//用例表
type Case struct {
    BaseModel   `bson:",inline"`
    CaseName string `bson:"caseName" json:"caseName"`
    // CaseSteps  e.g. [{"foo_obj_id": "api"}, {"var_obj_id": "testcase"}, {...}]
    CaseSteps interface{} `bson:"caseSteps" json:"caseSteps"`
    RelativePro []string `bson:"relativePro" json:"relativePro"`
    CaseDesc string `bson:"caseDesc" json:"caseDesc"`
    ConfigBaseUrl string `bson:"configBaseUrl" json:"configBaseUrl"`
    // ConfigVariables e.g.  [{"expected_status_code" : 200}, {...}]
    ConfigVariables interface{} `bson:"configVariables" json:"configVariables"`
    // ConfigOutput e.g.  ["session_token", ...]
    ConfigOutput []string `bson:"configOutput" json:"configOutput"`
    //ConfigVerify 默认为false
    ConfigVerify bool `bson:"configVerify" json:"configVerify"`
    //Extract e.g.  {"apiId": [{"extractName":"extractValue"}, {"session_token":"content.token"}],"apiId2": [{"extractName":"extractValue"}, {"session_token":"content.token"}]}
    Extract interface{} `bson:"extract" json:"extract"`
    //Validate e.g. {"apiId": [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]],"apiId2": [["eq", "status_code", "$expected_status_code"], ["eq", "content.headers.Host", "httpbin.org"]],
    Validate interface{} `bson:"validate" json:"validate"`
    //Variables e.g. {"apiId": [{"expected_status_code": 200}], "apiId2": [{"expected_status_code": 200}]}},
    Variables interface{} `bson:"variables" json:"variables"`
    //SetupHooks e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"],
    SetupHooks interface{} `bson:"setupHooks" json:"setupHooks"`
    //TeardownHooks e.g. ["${setup_hook_prepare_kwargs($request)}","${teardown_hook_sleep_N_secs($response, n_secs)}"]
    TeardownHooks interface{} `bson:"teardownHooks" json:"teardownHooks"`
}