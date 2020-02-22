package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
)

func main(){
    session, err := mgo.Dial("127.0.0.1:27017")
    mgo.SetDebug(true)
    if err != nil{
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    Conn := session.DB("testgo")
    log.Println("Connect Mongodb Success")
    var result interface{}
    err = Conn.C("project").Find(bson.M{"_id": "5e3d0a68268fcf112331ec8c"}).One(&result)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Printf("%+v", result)
}
