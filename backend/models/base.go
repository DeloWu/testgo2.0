package models

import (
    "gopkg.in/mgo.v2/bson"
)

//公共配置
type BaseModel struct {
    ID bson.ObjectId `bson:"_id" json:"id"`
    CreateTime int64 `bson:"createTime" json:"createTime"`
    UpdateTime int64 `bson:"updateTime" json:"updateTime"`
}

