package datasource

import (
    "gopkg.in/mgo.v2"
)

//IDb 数据库接口定义
type IDb interface {
    //Connect 初始化数据库配置
    Connect() error
    //DB 返回DB
    DB() *mgo.Database
    //Close() error
}
