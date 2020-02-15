package models

import (
    "gopkg.in/mgo.v2/bson"
    "time"
)

//公共配置
type BaseModel struct {
    ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
    CreateTime time.Time `json:"createTime"`
    UpdateTime time.Time `json:"updateTime"`
}

