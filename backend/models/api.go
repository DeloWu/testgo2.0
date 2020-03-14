package models

//接口表
type Api struct {
	BaseModel      `bson:",inline"`
	ApiName        string      `bson:"apiName" json:"apiName"`
	ApiBaseUrl     string      `bson:"apiBaseUrl" json:"apiBaseUrl"`
	ApiPath        string      `bson:"apiPath" json:"apiPath"`
	ApiMethod      string      `bson:"apiMethod" json:"apiMethod"`
	ApiContentType string      `bson:"apiContentType" json:"apiContentType"`
	ApiHeaders     interface{} `bson:"apiHeaders" json:"apiHeaders"`
	ApiParams      interface{} `bson:"apiParams" json:"apiParams"`
	ApiPostData    interface{} `bson:"apiPostData" json:"apiPostData"`
	ApiDesc        string      `bson:"apiDesc" json:"apiDesc"`
	RelativePro    []string    `bson:"relativePro" json:"relativePro"`
	// 格式为: [{"keyName":"", "value":""}, {, , }...]
	Extract interface{} `bson:"extract" json:"extract"`
	// 格式为: [{"keyName":"", "comparator":"", "value":""}, {, , }...]
	Validate interface{} `bson:"validate" json:"validate"`
	// 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
	Variables interface{} `bson:"variables" json:"variables"`
	//SetupHooks e.g. [{"value":"${setup_hook_prepare_kwargs($request)}"},{"value":"${teardown_hook_sleep_N_secs($response, n_secs)}"}],
	SetupHooks    interface{} `bson:"setupHooks" json:"setupHooks"`
	TeardownHooks interface{} `bson:"teardownHooks" json:"teardownHooks"`
	//文件的本地相对路径
	Location string `bson:"location" json:"location"`
}
