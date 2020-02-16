package models

import (
    "gopkg.in/mgo.v2/bson"
)

//公共配置
type BaseModel struct {
    ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
    CreateTime int64 `json:"createTime"`
    UpdateTime int64 `json:"updateTime"`
}

