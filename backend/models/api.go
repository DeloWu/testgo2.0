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
	// 格式为: [{"keyName":"", "valueType":"", "value":""}, {, , }...]
	Variables interface{} `bson:"variables" json:"variables"`
	// 格式为: [{"keyName":"", "value":""}, {, , }...]
	Extract interface{} `bson:"extract" json:"extract"`
	// 格式为: [{"keyName":"", "comparator":"", "value":""}], {, , }...
	Validate interface{} `bson:"validate" json:"validate"`
	//TODO: 前端组件做优化
	SetupHooks    []string `bson:"setupHooks" json:"setupHooks"`
	TeardownHooks []string `bson:"teardownHooks" json:"teardownHooks"`
}
