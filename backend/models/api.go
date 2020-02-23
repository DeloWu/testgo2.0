package models

import "gopkg.in/mgo.v2/bson"

//接口表
type Api struct {
    BaseModel `bson:",inline"`
    ApiName string `bson:"apiName" json:"apiName"`
    ApiUrl string `bson:"apiUrl" json:"apiUrl"`
    ApiMethod string `bson:"apiMethod" json:"apiMethod"`
    ApiContentType string `bson:"apiContentType" json:"apiContentType"`
    ApiHeader bson.M `bson:"apiHeader" json:"apiHeader"`
    ApiParams bson.M `bson:"apiParams" json:"apiParams"`
    ApiPostData bson.M `bson:"apiPostData" json:"apiPostData"`
    ApiDesc string `bson:"apiDesc" json:"apiDesc"`
    RelativePro []string `bson:"relativePro" json:"relativePro"`
    // Variables e.g.  [{"expected_status_code" : 200}, {...}]
    Variables interface{} `bson:"variables" json:"variables"`
    // Validate  e.g.  [["eq", "status_code", "$expected_status_code"], [...]]
    Validate interface{} `bson:"validate" json:"validate"`
}
