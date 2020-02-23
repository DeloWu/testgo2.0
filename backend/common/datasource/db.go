package datasource

import (
    "gopkg.in/mgo.v2"
    "log"
    "testgo/common/setting"

    //"gopkg.in/mgo.v2/bson"
)

type Db struct {
    Conn *mgo.Database
}

//Connect 初始化数据库配置
func(d *Db) Connect() error {
    conf := setting.Config.Database
    dbHost := conf.Host
    dbName := conf.Name
    session, err := mgo.Dial(dbHost)
    if err != nil{
        panic(err)
    }
    //session.SetMode(mgo.Monotonic, true)
    d.Conn = session.DB(dbName)
    d.Conn = session.DB("testgo")

    log.Println("Connect Mongodb Success")
    return nil
}

//DB 返回DB
func(d *Db) DB() *mgo.Database{
    return d.Conn
}

//关闭db
func(d *Db) Close() error{
    d.Conn.Session.Close()
    log.Println("Close Mongodb Success")
    return nil
}