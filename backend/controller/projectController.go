package controller

import (
    "github.com/astaxie/beego/validation"
    "gopkg.in/mgo.v2/bson"
    "log"

    //"log"
    "net/http"
    //"strconv"
    "testgo/common/codes"
    "testgo/models"

    //"github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    //"gopkg.in/mgo.v2/bson"
    "testgo/service"
    //debug
    "gopkg.in/mgo.v2"
    "testgo/common/datasource"

)

//依赖注入
type Project struct {
    Service service.IProjectService `inject:""`
}

//根据ID获取project
func (p *Project) GetProjectById(c *gin.Context) {
    valid := validation.Validation{}
    id := c.Query("id")
    valid.Required(id, "id")
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetProjectById(id)
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Project) GetProjectsByPagination(c *gin.Context) {
    valid := validation.Validation{}
    pageIndex, pageSize := GetPage(c)
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetProjectsByPagination(pageIndex, pageSize)
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)

}

//调试使用
func MockData(c *gin.Context) {
    var db datasource.Db
    var project models.Project
    log.Printf("%+v", db.Conn)
    //db.DB().C("project").Find(bson.M{}).One(&project)
    log.Printf("%+v", project)
    c.JSON(200, gin.H{
        "data": project,
    })
}

//debug
func FindDb(c *gin.Context) {
    log.Println("find db func start")
    var result models.Project
    session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil{
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    Conn := session.DB("testgo")
    log.Println("Connect Mongodb Success")
    err = Conn.C("project").Find(nil).One(&result)
    if err != nil{
        panic(err)
    }
    log.Println("find db finish, result is:")
    log.Printf("%+v", result)
    c.JSON(200, gin.H{
        "data": result,
    })
}

//debug
func FindDb1(c *gin.Context) {
    log.Println("find db func start")
    var result models.Project
    session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil{
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    Conn := session.DB("testgo")
    log.Println("Connect Mongodb Success")
    id := c.Query("id")
    err = Conn.C("project").FindId(bson.ObjectIdHex(id)).One(&result)
    if err != nil{
        panic(err)
    }
    log.Println("find db finish, result is:")
    log.Printf("%+v", result)
    c.JSON(200, gin.H{
        "data": result,
    })
}